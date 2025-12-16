package waconnect

import (
    "context"
    "fmt"
    "os"

    "go.mau.fi/whatsmeow"
    "go.mau.fi/whatsmeow/types"
    sqlstore "go.mau.fi/whatsmeow/store/sqlstore"
    _ "github.com/mattn/go-sqlite3"
    waLog "go.mau.fi/whatsmeow/util/log"

    whatsgotypes "github.com/WinterSunset95/WhatsGo/WhatsGoTypes"

    "github.com/mdp/qrterminal/v3"
)

var (
	WAClient        *whatsmeow.Client
	CurrentChat     types.JID
	WhatsGoDatabase whatsgotypes.Database
	DBPath          string
)



func WAConnect(whatsGoDb string) (*whatsmeow.Client, error) {
    if WhatsGoDatabase == nil {
	    WhatsGoDatabase = make(whatsgotypes.Database)
    }
    logger := waLog.Stdout("WhatsGo", "INFO", true)

    ctx := context.Background()

    // sqlite store
    container, err := sqlstore.New(ctx, "sqlite3", "file:" + whatsGoDb + "?_foreign_keys=on", logger)
    if err != nil {
        return nil, fmt.Errorf("sqlstore.New: %w", err)
    }

    device, err := container.GetFirstDevice(ctx)
    if err != nil {
        return nil, fmt.Errorf("GetFirstDevice: %w", err)
    }

    client := whatsmeow.NewClient(device, logger)

    // Si no está logueado, QR pairing
    if client.Store.ID == nil {
        ctx := context.Background()

        qrChan, err := client.GetQRChannel(ctx)
        if err != nil {
            return nil, fmt.Errorf("GetQRChannel: %w", err)
        }

        if err := client.Connect(); err != nil {
            return nil, fmt.Errorf("Connect: %w", err)
        }

        for evt := range qrChan {
            // evt.Event suele ser "code" o "success", según versión
            fmt.Printf("QR event: %#v\n", evt)
            if evt.Event == "code" {
                fmt.Println("Scan this QR (terminal):", evt.Code)
	        qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
            } else if evt.Event == "success" {
                fmt.Println("Paired successfully")
                break
            } else if evt.Event == "timeout" {
                return nil, fmt.Errorf("QR timeout")
            } else if evt.Event == "err-client-outdated" {
                return nil, fmt.Errorf("client outdated according to server")
            }
        }
    } else {
        // Ya hay sesión guardada
        if err := client.Connect(); err != nil {
            return nil, fmt.Errorf("Connect: %w", err)
        }
    }

    return client, nil
}

