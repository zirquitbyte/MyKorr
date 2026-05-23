package fake

import (
	"fmt"
	"time"
	"math/rand"

	"MyKorr/internal/db"
	"MyKorr/internal/models"
)

type FakeConnector struct {}

func NewFakeConnector() *FakeConnector {
	rand.Seed(time.Now().UnixNano())
    return &FakeConnector{}
}

// message store interface instead of specific DB
func (c *FakeConnector) Start(store db.MessageStore) {
	messages := []string{
		"So, uhh, you're saying you slept ONLY %d HOURS last night?? Well, I guess it's better than %d...",
		"Meet me at the usual place, in %d minutes? Or are you gonna be %d hours late again?",
		"Statistically speaking, %d percent of missing knights are caused by %d percent of dragons.",
		"I heard IKEA has restocked BLÅHAJ and DJUNGELSKOG! I'm buying at least %d! I can get you some, would %d be enough?",
		"No, I really think %d is smaller than %d. I mean, maybe not necessarily in a literal way...",
		"Ugh, every time I do my laundry, there's like %d more socks I can't find. That's gotta be like, at least %d sock thiefs at large!",
		"Listen. Listen. All I'm saying is. If %d people sent me a cent. I'd have... More than %d cents. It's like, literally math.",
		"Quick, quick, what happened in year %d? I'm in my protohistory exam and panopticon is looking the other way. Hurry up, I only have %d minutes.",
		"%d days till release? Can we maybe make it %d?",
		"Yesterday I had (I checked!) %d hamsters. This morning I woke up to %d...",
		"I think I'll finish this project in about %d energy drink cans. That sound reasonable? I already went through %d crates btw.",
		"No, we can not adopt %d new cats. No, not even %d. Like, how does that even make sense?",
		"I'm so tired of having like %d messenger apps on my phone. %d of them I don't even use. Someone should do something about that.",
    	}
		for {
		a:= rand.Intn(13) + 8
		b:= rand.Intn(4) + 2

		body := fmt.Sprintf(
			messages[rand.Intn(len(messages))],
			a, b,
		)

        msg := models.Message{
			Timestamp: time.Now(),
			ConversationID: "1",
			Source: "fake",
			Sender: "ThatOneFriend",
			Body: body,
			Attachments: true,
			Reply_to: "0",
        }

        store.Save(msg)

		fmt.Println("[" + msg.Timestamp.Format("15:04")+ "]", msg.Sender, ":", msg.Body,)

        time.Sleep(time.Duration(rand.Intn(6)+2) * time.Second)
    }
}

