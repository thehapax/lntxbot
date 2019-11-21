package t

type Key string

const (
	NO         Key = "No"
	YES            = "Yes"
	CANCEL         = "Cancel"
	CANCELED       = "Canceled"
	COMPLETED      = "Completed"
	CONFIRM        = "Confirm"
	FAILURE        = "Failure"
	PROCESSING     = "Processing"
	WITHDRAW       = "Withdraw"
	ERROR          = "Error"
	CHECKING       = "Checking"
	TXPENDING      = "TxPending"
	TXCANCELED     = "TxCanceled"
	UNEXPECTED     = "Unexpected"

	CLAIMFAILED = "ClaimFailed"

	CALLBACKCOINFLIPWINNER = "CallbackCoinflipWinner"
	CALLBACKWINNER         = "CallbackWinner"
	CALLBACKERROR          = "CallbackError"
	CALLBACKEXPIRED        = "CallbackExpired"
	CALLBACKATTEMPT        = "CallbackAttempt"
	CALLBACKSENDING        = "CallbackSending"

	INLINEINVOICERESULT  = "InlineInvoiceResult"
	INLINEGIVEAWAYRESULT = "InlineGiveawayResult"
	INLINEGIVEFLIPRESULT = "InlineGiveflipResult"
	INLINECOINFLIPRESULT = "InlineCoinflipResult"
	INLINEHIDDENRESULT   = "InlineHiddenResult"

	LNURLUNSUPPORTED = "LnurlUnsupported"
	LNURLAUTHSUCCESS = "LnurlAuthSuccess"
	LNURLPAYPROMPT   = "LnurlPayPrompt"

	USERALLOWED       = "UserAllowed"
	SPAMFILTERMESSAGE = "SpamFilterMessage"

	PAYMENTFAILED       = "PaymentFailed"
	PAIDMESSAGE         = "PaidMessage"
	DBERROR             = "DBError"
	INSUFFICIENTBALANCE = "InsufficientBalance"

	PAYMENTRECEIVED      = "PaymentReceived"
	FAILEDTOSAVERECEIVED = "FailedToSaveReceived"

	SPAMMYMSG           = "SpammyMsg"
	COINFLIPSENABLEDMSG = "CoinflipsEnabledMsg"
	LANGUAGEMSG         = "LanguageMsg"
	TICKETMSG           = "TicketMsg"
	FREEJOIN            = "FreeJoin"

	HELPINTRO   = "HelpIntro"
	HELPSIMILAR = "HelpSimilar"
	HELPMETHOD  = "HelpMethod"

	RECEIVEHELP = "receiveHelp"

	PAYHELP = "payHelp"

	SENDHELP = "sendHelp"

	TRANSACTIONSHELP = "transactionsHelp"

	BALANCEHELP = "balanceHelp"

	GIVEAWAYHELP            = "giveawayHelp"
	GIVEAWAYMSG             = "GiveAwayMsg"
	GIVEAWAYCLAIM           = "GiveAwayClaim"
	GIVEAWAYSATSGIVENPUBLIC = "GiveawaySatsGivenPublic"

	COINFLIPHELP      = "coinflipHelp"
	COINFLIPWINNERMSG = "CoinflipWinnerMsg"
	COINFLIPGIVERMSG  = "CoinflipGiverMsg"
	COINFLIPAD        = "CoinflipAd"
	COINFLIPJOIN      = "CoinflipJoin"
	COINFLIPOVERQUOTA = "CoinflipOverQuota"
	COINFLIPRATELIMIT = "CoinflipRateLimit"

	GIVEFLIPHELP      = "giveflipHelp"
	GIVEFLIPMSG       = "GiveFlipMsg"
	GIVEFLIPWINNERMSG = "GiveflipWinnerMsg"
	GIVEFLIPAD        = "GiveflipAd"
	GIVEFLIPJOIN      = "GiveflipJoin"

	FUNDRAISEHELP        = "fundraiseHelp"
	FUNDRAISEAD          = "FundraiseAd"
	FUNDRAISEJOIN        = "FundraiseJoin"
	FUNDRAISECOMPLETE    = "FundraiseComplete"
	FUNDRAISERECEIVERMSG = "FundraiseReceiverMsg"
	FUNDRAISEGIVERMSG    = "FundraiseGiverMsg"

	BLUEWALLETHELP         = "bluewalletHelp"
	BLUEWALLETCREDENTIALS  = "BluewalletCredentials"
	APIPASSWORDUPDATEERROR = "APIPasswordUpdateError"
	APICREDENTIALS         = "APICredentials"

	HIDEHELP             = "hideHelp"
	REVEALHELP           = "revealHelp"
	HIDDENREVEALBUTTON   = "HiddenRevealButton"
	HIDDENDEFAULTPREVIEW = "HiddenDefaultPreview"
	HIDDENWITHID         = "HiddenWithId"
	HIDDENSOURCEMSG      = "HiddenSourceMsg"
	HIDDENREVEALMSG      = "HiddenRevealMsg"
	HIDDENMSGNOTFOUND    = "HiddenMsgNotFound"
	HIDDENSHAREBTN       = "HiddenShareBtn"

	BITFLASHHELP         = "bitflashHelp"
	BITFLASHCONFIRM      = "BitflashConfirm"
	BITFLASHTXQUEUED     = "BitflashTxQueued"
	BITFLASHFAILEDTOSAVE = "BitflashFailedToSave"
	BITFLASHLIST         = "BitflashList"

	MICROBETHELP                = "microbetHelp"
	MICROBETBETHEADER           = "MicrobetBetHeader"
	MICROBETPAIDBUTNOTCONFIRMED = "MicrobetPaidButNotConfirmed"
	MICROBETPLACING             = "MicrobetPlacing"
	MICROBETPLACED              = "MicrobetPlaced"
	MICROBETLIST                = "MicrobetList"
	MICROBETBALANCE             = "MicrobetBalance"

	BITREFILLHELP            = "bitrefillHelp"
	BITREFILLINVENTORYHEADER = "BitrefillInventoryHeader"
	BITREFILLPACKAGESHEADER  = "BitrefillPackagelHeader"
	BITREFILLNOPROVIDERS     = "BitrefillNoProviders"
	BITREFILLCONFIRMATION    = "BitrefillConfirmation"
	BITREFILLFAILEDSAVE      = "BitrefillFailedToSave"
	BITREFILLPURCHASEDONE    = "BitrefillPurchaseDone"
	BITREFILLPURCHASEFAILED  = "BitrefillPurchaseFailed"
	BITREFILLCOUNTRYSET      = "BitrefillCountrySet"
	BITREFILLINVALIDCOUNTRY  = "BitrefillInvalidCountry"

	SATELLITEHELP              = "satelliteHelp"
	SATELLITEFAILEDTOSTORE     = "SatelliteFailedToStore"
	SATELLITEFAILEDTOGET       = "SatelliteFailedToGet"
	SATELLITEPAID              = "SatellitePaid"
	SATELLITEFAILEDTOPAY       = "SatelliteFailedToPay"
	SATELLITETRANSMISSIONERROR = "SatelliteTransmissionError"
	SATELLITELIST              = "SatelliteList"

	FUNDBTCHELP   = "fundbtcHelp"
	FUNDBTCFAIL   = "fundbtcFail"
	FUNDBTCFINISH = "fundbtcFinish"

	BITCLOUDSHELP           = "bitcloudsHelp"
	BITCLOUDSCREATEHEADER   = "BitcloudsCreateHeader"
	BITCLOUDSCREATED        = "BitcloudsCreated"
	BITCLOUDSSTOPPEDWAITING = "BitcloudsStoppedWaiting"
	BITCLOUDSNOHOSTS        = "BitcloudsNoHosts"
	BITCLOUDSHOSTSHEADER    = "BitcloudsHostsHeader"
	BITCLOUDSSTATUS         = "BitcloudsStatus"
	BITCLOUDSREMINDER       = "BitcloudsReminder"

	QIWIHELP             = "qiwiHelp"
	YANDEXHELP           = "yandexHelp"
	LNTORUBCONFIRMATION  = "LNToRubConfirmation"
	LNTORUBFULFILLED     = "LNToRubFulfilled"
	LNTORUBCANCELED      = "LNToRubCanceled"
	LNTORUBMISSINGTARGET = "LNToRubMissingTarget"
	LNTORUBFIATERROR     = "LNToRubFiatError"
	LNTORUBORDERLIST     = "LNToRubOrderList"
	LNTORUBDEFAULTTARGET = "LNToRubDefaultTarget"

	GIFTSHELP       = "giftsHelp"
	GIFTSCREATED    = "GiftsCreated"
	GIFTSFAILEDSAVE = "GiftsFailedSave"
	GIFTSLIST       = "GiftsList"
	GIFTSSPENTEVENT = "GiftsSpentEvent"

	POKERHELP         = "pokerHelp"
	POKERDEPOSITFAIL  = "PokerDepositFail"
	POKERWITHDRAWFAIL = "PokerWithdrawFail"
	POKERSTATUS       = "PokerStatus"
	POKERNOTIFY       = "PokerNotify"
	POKERNOTIFYFRIEND = "PokerNotifyFriend"
	POKERSUBSCRIBED   = "PokerSubscribed"
	POKERSECRETURL    = "PokerSecretURL"
	POKERBALANCE      = "PokerBalance"

	SATS4ADSHELP       = "sats4adsHelp"
	SATS4ADSTOGGLE     = "Sats4adsToggle"
	SATS4ADSBROADCAST  = "Sats4adsBroadcast"
	SATS4ADSPRICETABLE = "Sats4adsPriceTable"
	SATS4ADSADFOOTER   = "Sats4adsAdFooter"

	PAYWALLHELP      = "paywallHelp"
	PAYWALLBALANCE   = "PaywallBalance"
	PAYWALLCREATED   = "PaywallCreated"
	PAYWALLLISTLINKS = "PaywallListLinks"
	PAYWALLPAIDEVENT = "PaywallPaidEvent"

	ETLENEUMFAILEDTOPAY = "EtleneumFailedToPay"

	TOGGLEHELP = "toggleHelp"

	HELPHELP = "helpHelp"

	STOPHELP = "stopHelp"

	CONFIRMINVOICE     = "ConfirmInvoice"
	FAILEDDECODE       = "FailedDecode"
	NOINVOICE          = "NoInvoice"
	BALANCEMSG         = "BalanceMsg"
	TAGGEDBALANCEMSG   = "TaggedBalanceMsg"
	FAILEDUSER         = "FailedUser"
	LOTTERYMSG         = "LotteryMsg"
	INVALIDPARTNUMBER  = "InvalidPartNumber"
	INVALIDAMOUNT      = "InvalidAmount"
	USERSENTTOUSER     = "UserSentToUser"
	USERSENTYOUSATS    = "UserSentYouSats"
	RECEIVEDSATSANON   = "ReceivedSatsAnon"
	FAILEDSEND         = "FailedSend"
	QRCODEFAIL         = "QRCodeFail"
	SAVERECEIVERFAIL   = "SaveReceiverFail"
	CANTSENDNORECEIVER = "CantSendNoReceiver"
	GIVERCANTJOIN      = "GiverCantJoin"
	CANTJOINTWICE      = "CantJoinTwice"
	CANTCANCEL         = "CantCancel"
	FAILEDINVOICE      = "FailedInvoice"
	ZEROAMOUNTINVOICE  = "ZeroAmountInvoice"
	INVALIDAMT         = "InvalidAmt"
	STOPNOTIFY         = "StopNotify"
	WELCOME            = "Welcome"
	WRONGCOMMAND       = "WrongCommand"
	RETRACTQUESTION    = "RetractQuestion"
	RECHECKPENDING     = "RecheckPending"
	TXNOTFOUND         = "TxNotFound"
	TXINFO             = "TxInfo"
	TXLIST             = "TxList"

	TUTORIALWALLET = "TutorialWallet"
	TUTORIALBLUE   = "TutorialBlue"
	TUTORIALAPPS   = "TutorialApps"
)