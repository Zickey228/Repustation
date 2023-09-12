package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"strconv"
)

const deadAdd = "1AkFuweFVhr4pVkGnoua16qWtJpUb8NhZh"

func GUI() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Repu-wallet")

	//TAB 1
	sender := widget.NewEntry()
	reviewTarget := widget.NewEntry()
	review := widget.NewMultiLineEntry()
	investment := widget.NewEntry()
	polarity := widget.NewEntry()

	submit := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Sender", Widget: sender}, {Text: "Target", Widget: reviewTarget}, {Text: "Investment", Widget: investment}, {Text: "Polarity", Widget: polarity}, {Text: "Review", Widget: review}},
		OnSubmit: func() {
			//submit logic
			//log.Println("Form submitted:", entry.Text)
			//log.Println("multiline:", textArea.Text)
			//submitReivew("1", entry1.Text, textArea1.Text)
			//发送investment给销毁地址
			//investAmount, _ := strconv.Atoi(investment.Text)
			//send(sender.Text, deadAdd, investAmount)
			submitReview(sender.Text, reviewTarget.Text, review.Text, investment.Text, polarity.Text)
			dialog.ShowInformation("Congrats", "Reviews is submitted", myWindow)
		},
	}
	label1 := widget.NewLabel("Wallet address:")
	walletText := ""
	walletAdd := widget.NewLabel(walletText)
	//label2 := widget.NewLabel("Private key:")
	//privateKey := widget.NewLabel("N/A")

	//label3 := widget.NewLabel("Nickname :")
	//nickname := widget.NewLabel("N/A")
	//创建私钥
	register := widget.NewLabel("Press button to create new wallet:")
	//保存私钥并创建
	//fileEntry := widget.NewEntry()
	registerAndSave := widget.NewButton("generate & save", func() {

		//wallet, priKey, pubKey := NewWallet()
		//add := string(wallet.GetAddress())

		//walletAdd.SetText(add)
		//publicKey.SetText(public)
		//privateKey.SetText(priKey.D.String())
		//nickname.SetText(fileEntry.Text)
		//fmt.Printf("pri:%x,pub:%x,wallet:%x", priKey.D, pubKey, add)
		//分别保存私钥，公钥，钱包地址
		//SaveToFile(priKey.D.String(), fileEntry.Text+"Pri")
		//SaveToFile(string(pubKey), fileEntry.Text+"Pub")
		//SaveToFile(add, fileEntry.Text+"Wal")
		//wallets.SaveToFile(address, fileEntry.Text)

		wallets, _ := NewWallets()
		address := wallets.CreateWallet()
		wallets.SaveToFile()
		walletAdd.SetText(address)
		fmt.Printf("你的新钱包地址是: %s\n", address)
		dialog.ShowInformation("Success", address, myWindow)
	})
	//registerAndSave := widget.NewButton("generate & save", func() {
	//	add.SetText("789")
	//	newFileName := "./wallet" + "789"
	//	file, err := os.Create(newFileName)
	//	if err != nil {
	//		fmt.Printf(err.Error())
	//		return
	//	}
	//	defer file.Close()
	//	file.WriteString("789")
	//	dialog.ShowInformation("congrats", "export wallet success", myWindow)
	//})

	registerContainer := widget.NewVBox(register, registerAndSave)
	infoContainer := widget.NewVBox(label1, walletAdd)

	//登录容器
	loginLable := widget.NewLabel("List all wallets:")
	//loginEntry := widget.NewEntry()
	loginBtn := widget.NewButton("List wallet addresses", func() {
		//priData, err := LoadFileToHex(loginEntry.Text + "Pri")
		//pubData, err := LoadFileToHex(loginEntry.Text + "Pub")
		//walData, err := LoadFileToHex(loginEntry.Text + "Wal")
		//if err != nil {
		//	dialog.ShowInformation("Failed", err.Error(), myWindow)
		//}
		//private, _ := new(big.Int).SetString(string(priData), 10)
		//fmt.Printf("prikey:%x", private)
		//fmt.Printf("pubkey:%x", pubData)
		//fmt.Printf("walAdd:%x", walData)
		////privateKey.SetText(private.String())
		////nickname.SetText(loginEntry.Text)
		//walletAdd.SetText(string(walData))
		adds := listWallets()
		for _, s := range adds {
			walletText += s + "\n"
		}
		walletAdd.SetText(walletText)
	})
	loginContainer := widget.NewVBox(loginLable, loginBtn)
	//login := &widget.Form{
	//	Items: []*widget.FormItem{ // we can specify items in the constructor
	//		{Text: "Login", Widget: entry}, {Text: "Review", Widget: textArea}},
	//	OnSubmit: func() {
	//		//submit logic
	//		//log.Println("Form submitted:", entry.Text)
	//		//log.Println("multiline:", textArea.Text)
	//		submitReivew("1", entry.Text, textArea.Text)
	//		dialog.ShowInformation("Congrats", "You are using following wallet", myWindow)
	//	},
	//}

	//检查评论容器
	checkLable1 := widget.NewLabel("Enter the target")
	checkEntry1 := widget.NewEntry()
	checkBtn1 := widget.NewButton("Show reviews", func() {
		//从网站请求评论列表
		fmt.Printf("Reviews")
	})
	checkBtn2 := widget.NewButton("Check reviews", func() {
		//比对网站和区块链评论
		if checkReview(checkEntry1.Text) {
			dialog.ShowInformation("Congrats", "Reviews is integrated", myWindow)
		} else {
			dialog.ShowInformation("Sorry", "Someone tampered reviews", myWindow)
		}
		//fmt.Printf("Check reviews")
	})
	checkReviewContainer := widget.NewVBox(checkLable1, checkEntry1, checkBtn1, checkBtn2)

	//交易容器
	senderWallet := widget.NewLabel("Sender address")
	senderEntry := widget.NewEntry()
	target := widget.NewLabel("Target")
	targetEntry := widget.NewEntry()
	amount := widget.NewLabel("Amount")
	amountEntry := widget.NewEntry()

	txBtn1 := widget.NewButton("send", func() {
		//发送代币逻辑
		amount, _ := strconv.Atoi(amountEntry.Text)
		send(senderEntry.Text, targetEntry.Text, amount)
		fmt.Printf("send tokens")
	})
	txLabel2 := widget.NewLabel("Balance:")
	balance := widget.NewLabel("N/A")
	balEntry := widget.NewEntry()
	checkBalance := widget.NewButton("Check balance", func() {
		//fmt.Printf("check balance")
		addr := balEntry.Text
		fmt.Printf("opening %x and requesting balance", addr)
		bal := balanceReq(addr)
		balance.SetText(bal)
	})
	transactionContainer := widget.NewVBox(senderWallet, senderEntry, target, targetEntry, amount, amountEntry, txBtn1, balance, txLabel2, balance, balEntry, checkBalance)

	//rating容器

	ratingLabel1 := widget.NewLabel("Sender address:")
	ratingEntry1 := widget.NewEntry()
	ratingLabel2 := widget.NewLabel("Receiver address:")
	ratingEntry2 := widget.NewEntry()
	ratingLabel3 := widget.NewLabel("Review ID:")
	ratingEntry3 := widget.NewEntry()
	ratingLabel4 := widget.NewLabel("Polarity:")
	ratingEntry4 := widget.NewEntry()
	ratingBtn := widget.NewButton("sending rating to full node", func() {
		//评分逻辑
		sendRating(ratingEntry1.Text, ratingEntry2.Text, ratingEntry3.Text, ratingEntry4.Text)
	})
	ratingContainer := widget.NewVBox(ratingLabel1, ratingEntry1, ratingLabel2, ratingEntry2, ratingLabel3, ratingEntry3, ratingLabel4, ratingEntry4, ratingBtn)

	//同步容器
	syncBtn1 := widget.NewButton("Sync. with full node", func() {
		fmt.Printf("snynchronizing with full node......")
		recv_file(dns)
		dialog.ShowInformation("Successful", "Congrats, light node succesfully sychronised from peer", myWindow)
	})
	syncContainer := widget.NewVBox(syncBtn1)

	//所有标签
	tabs := container.NewAppTabs(
		container.NewTabItem("INFO", infoContainer),
		container.NewTabItem("Submit reviews", submit),
		container.NewTabItem("Check reviews", checkReviewContainer),
		container.NewTabItem("transaction", transactionContainer),
		container.NewTabItem("Register", registerContainer),
		container.NewTabItem("List addresses", loginContainer),
		container.NewTabItem("Rating", ratingContainer),
		container.NewTabItem("Sync.", syncContainer),
		//container.NewTabItem("Login", login),
		//container.NewTabItem("Sync",sync)
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("ReadMe: You can use this application to \n 1. Submit reviews in a democratic way \n 2. Earn tokens \n 3. Trade tokens")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)
	myWindow.Resize(fyne.Size{1000, 600})
	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
