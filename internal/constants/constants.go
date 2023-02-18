package constants

import (
	"bitbucket.org/esportsph/minigame-backend-golang/internal/settings"
	"bitbucket.org/esportsph/minigame-backend-golang/internal/utils"
)

type CURRENCY_TYPE = float64

const MAX_FUTURE_EVENTS = 16
const LOL_TOWER_GAME_DURATION = 14
const LOL_TOWER_GAME_ID = 5
const LOL_TOWER_TABLE_11 = 11
const LOL_TOWER_SKIP_COUNT = 3
const LOL_TOWER_DEFAULT_LEVEL = 1
const LOL_TOWER_BETTING_DURATION = 7
const LOL_TOWER_STOP_BETTING_DURATION = 3
const LOL_TOWER_SHOW_RESULT_DURATION = 4
const MG_DEFAULT_DB_STATUS = "db"
const LOL_TOWER_ES_GAME_ID = 33
const LOL_TOWER_GAME_NAME = "lol-tower"
const LOL_TOWER_COMPETITION = "Lootbox"
const TRANSACTION_TYPE_BET = "bet"
const LOL_TOWER_MARKET_OPTION = "match"
const LOL_TOWER_MARKET_TYPE = 18
const BET_TYPE_SPOR = "SPOR"

const LOL_TOWER_MAX_LEVEL = 10
const LOL_SKIP_SELECTION = "s"
const LOL_PAYOUT_SELECTION = "p"
const CUSTOM_WIN_SELECTION = "w"
const CUSTOM_LOSS_SELECTION = "l"
const DEFAULT_MEMBER_ODDS_STYLE = "hongkong"
const MEMBER_CODE_MASK_COUNT = 4
const DEFAULT_GROUND_TYPE = "Neutral"
const MAX_WS_RETRIES = 3

var GAME_MARKET_NAME = map[int]string{
	LOL_TOWER_GAME_ID: "Tower",
}
var LOL_BET_SELECTION = []string{"1", "2", "3", "4", "5"}
var LOL_EXTRA_SELECTION = []string{LOL_SKIP_SELECTION, LOL_PAYOUT_SELECTION}

var LOL_TOWER_TABLES = [...]int64{
	LOL_TOWER_TABLE_11,
}

// all game tables
var TABLES = [...]int64{
	LOL_TOWER_TABLE_11,
}

func init() {
	if utils.InArray(settings.ENVIRONMENT, []string{"local", "dev"}) {
		LOL_EXTRA_SELECTION = append(LOL_EXTRA_SELECTION, CUSTOM_WIN_SELECTION)
		LOL_EXTRA_SELECTION = append(LOL_EXTRA_SELECTION, CUSTOM_LOSS_SELECTION)
	}
}

// default paginations
const PAGINATION_DEFAULT_LIMIT_PER_PAGE = 1000
const PAGINATION_DEFAULT_PAGE = 1
const PAGINATION_SORT_FEILD = "ctime asc"

//new implementation
const (
	TicketTypeDB     = "db"
	TicketTypeInplay = "inplay"
)

const (
	OddsStyleEuro     = "euro"
	OddsStyleHongkong = "hongkong"
	OddsStyleMalay    = "malay"
	OddsStyleIndo     = "indo"
)

const (
	HashStatusQueued = "queued"
	HashStatusActive = "active"
	HashStatusDone   = "done"
)

const (
	EVENT_STATUS_ENABLED = iota
	EVENT_STATUS_ACTIVE
	EVENT_STATUS_DISABLED
	EVENT_STATUS_FOR_SETTLEMENT
	EVENT_STATUS_SETTLEMENT_IN_PROGRESS
	EVENT_STATUS_SETTLED
	EVENT_CANCELLED
)

const (
	SETTLEMENT_STATUS_RESULTING = "resulting"
	SETTLEMENT_STATUS_SETTLED   = "settled"
	SETTLEMENT_STATUS_OPEN      = "open"
	SETTLEMENT_STATUS_CONFIRMED = "confirmed"
)

const (
	TICKET_STATUS_PLACED = iota
	TICKET_STATUS_PAYMENT_CONFIRMATION
	TICKET_STATUS_PAYMENT_CONFIRMED
	TICKET_STATUS_PAYMENT_FAILED
	TICKET_STATUS_SETTLEMENT_IN_PROGRESS
	TICKET_STATUS_SETTLED_PENDING_PAYOUT
	TICKET_STATUS_SETTLED_PAYOUT_IN_PROGRESS
	TICKET_STATUS_SETTLED
	TICKET_STATUS_CANCELLED_PENDING_PAYOUT
	TICKET_STATUS_CANCELLED_PAYOUT_IN_PROGRESS
	TICKET_STATUS_CANCELLED
	TICKET_STATUS_SETTLED_WAITING_FOR_EVENT_SETTLEMENT
)

const (
	TICKET_RESULT_WIN = iota
	TICKET_RESULT_LOSS
)
