my fork of [wintersunset95's WhatsGo](https://github.com/WinterSunset95/WhatsGo)

# WhatsGo
* A command line whatsapp client
![whatsgo](./whatsgo.png)

### Core Features
* ~~View image messages~~ 
* ~~View video messages~~
* ~~View sticker messages~~ 
* ~~View documents~~
* Search for contacts 
* Recieve read, sent and delivered status 
* Send message in a group 
* Send photos
* Send videos
* Send documents

### Planned Features
* fix what's crossed above

## Requirements
* Go 1.25
* feh (for viewing images)
* mpv (for viewing videos)

## Installation
#### Clone and run

```
git clone https://github.com/poweredgeR415/WhatsGo
``` 

```
cd WhatsGo
```

```
go run .
```
#### Optionally, you can just run the pre-built binary
```
./WhatsGo
```


## Usage
### first run
* a qr code will print on the terminal to authenticate with your account
* syncing of messages will start, you'll likely need to exit out and open again to see the initial fetched list

### general usage
* There are four main sections in the program:
    * Search: Search for contacts
    * Contacts: List of contacts, will filter based on 'Search'. Arrow keys to navigate, Enter to select.
    * Chat: A list of messages with the selected contact. Arrow keys to navigate, Enter to select.
    * Message: Type your message here. Press Enter to send.
* On running the program, you'll be on the 'Search' section.
* Use the Tab key to switch between sections.
* On the 'Chat' section, you can press enter on a media message (sticker, video, image) to view it.

## Important Notes
* The program often breaks on the first run **yes**
* Images and videos are downloaded in the background *to directory ~/.whatsgo*. It might take a while before you can see them.
* This is my FIRST golang project and I am basically bullshitting my way through. *me too bro but that's the beauty of it*
