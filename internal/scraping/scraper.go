package scraping

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

type Salat struct {
	Name string
	Hour time.Time
}

type errorScraping struct {
	err     error
	message string
}

const bouzignacMasjid = "https://mawaqit.net/en/mosquee-de-bouzignac-tours-37000-france-1"

func GetSalatTime() [5]Salat {
	//Setup headless browser
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	var nodes []*cdp.Node

	errorRunBrowser := errorScraping{
		message: "Erreur lors du lancement du browser sur le lien : ",
	}
	errorRunBrowser.err = chromedp.Run(ctx,
		chromedp.Navigate(bouzignacMasjid),
		chromedp.Nodes(".prayers > div", &nodes, chromedp.ByQueryAll),
	)

	displayErrorConsole(errorRunBrowser)

	return scrapingMawaqitWebsite(ctx, nodes)
}

func scrapingMawaqitWebsite(ctx context.Context, nodes []*cdp.Node) [5]Salat {
	var salatToday [5]Salat
	for salatNumber, node := range nodes {
		var salatName, salatTimeStr string

		errNodeScraping := errorScraping{
			message: "Erreur lors de la récupération des noeuds css : ",
		}
		errNodeScraping.err = chromedp.Run(ctx,
			chromedp.Text("div.name", &salatName, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text("div.time div", &salatTimeStr, chromedp.ByQuery, chromedp.FromNode(node)),
		)

		if salatName == "" || salatTimeStr == "" {
			log.Fatalf("Le nom de la prière => %s, ou l'heure est vide => %s", salatName, salatTimeStr)
		}

		displayErrorConsole(errNodeScraping)

		errParseDate := errorScraping{
			message: "Erreur lors de la conversion de l'heure : ",
		}
		var salatTime time.Time
		salatTime, errParseDate.err = time.Parse("15:04", salatTimeStr)
		displayErrorConsole(errParseDate)

		aSalat := Salat{
			Name: salatName,
			Hour: salatTime,
		}

		fmt.Println(aSalat)
		salatToday[salatNumber] = aSalat
	}
	return salatToday
}

func displayErrorConsole(errScraping errorScraping) {
	if errScraping.err != nil {
		log.Fatal(errScraping.message, errScraping.err)
	}
}
