package adapters

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"wallet/source/domain"
	"wallet/source/domain/messages"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	gofrsuuid "github.com/gofrs/uuid"
	"github.com/google/uuid"
)

type EventstoreWalletRepository struct {
	EventstoreClient *esdb.Client
}

func (r *EventstoreWalletRepository) Add(wallet *domain.Wallet) error {
	walletEvents := wallet.GetEvents()
	outputEvents := make([]esdb.EventData, len(walletEvents))

	for i, event := range walletEvents {
		data, err := json.Marshal(event)
		if err != nil {
			log.Println("error encoding the message body:", err)
			return err
		}

		// Eventstore requires gofrs.UUID
		eventId, err := gofrsuuid.FromString(event.GetKey())
		if err != nil {
			log.Println("error creating the event id:", err)
			return err
		}

		outputEvents[i] = esdb.EventData{
			EventID:     eventId,
			EventType:   event.GetName(),
			ContentType: esdb.JsonContentType,
			Data:        data,
		}
	}

	_, err := r.EventstoreClient.AppendToStream(
		context.Background(),
		fmt.Sprint("wallet-", wallet.GetId().String()),
		esdb.AppendToStreamOptions{},
		outputEvents...,
	)

	if err != nil {
		log.Println("error appending to stream:", err)
		return err
	}

	return nil
}

func (r *EventstoreWalletRepository) Get(id uuid.UUID) (*domain.Wallet, error) {
	options := esdb.ReadStreamOptions{
		From:      esdb.Start{},
		Direction: esdb.Forwards,
	}

	stream, err := r.EventstoreClient.ReadStream(
		context.Background(),
		fmt.Sprint("wallet-", id.String()),
		options,
		100,
	)

	if err != nil {
		log.Println("error reading stream:", err)
		return nil, err
	}

	defer stream.Close()

	walletEvents := make([]messages.Message, 0)

	for {
		inputEvent, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Println("error reading from stream:", err)
			return nil, err
		}

		// Decodes event data into a domain event the implements Message interface
		decodeEvent := func(walletEvent messages.Message) error {
			if err := json.Unmarshal(inputEvent.Event.Data, &walletEvent); err != nil {
				log.Println("error decoding the message body:", err)
				return err
			}
			walletEvents = append(walletEvents, walletEvent)
			return nil
		}

		switch inputEvent.Event.EventType {
		case "WalletCreated":
			var walletEvent messages.WalletCreated
			if err := decodeEvent(&walletEvent); err != nil {
				return nil, err
			}
		case "DepositReceived":
			var walletEvent messages.DepositReceived
			if err := decodeEvent(&walletEvent); err != nil {
				return nil, err
			}
		}
	}

	wallet := domain.NewFromEvents(walletEvents)

	log.Println("wallet loaded:", wallet.GetId())

	return wallet, nil
}
