package record

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: Record
const MajorVersion = 1
const MinorVersion = 13

var _ext *x.Extension

func Ext() *x.Extension {
	return _ext
}

// simple ('xcb', 'Record', 'CONTEXT')
type Context uint32

// simple ('xcb', 'Record', 'ElementHeader')
type ElementHeader uint8

// enum HType
const (
	HTypeFromServerTime     = 1
	HTypeFromClientTime     = 2
	HTypeFromClientSequence = 4
)

// simple ('xcb', 'Record', 'ClientSpec')
type ClientSpec uint32

// enum CS
const (
	CSCurrentClients = 1
	CSFutureClients  = 2
	CSAllClients     = 3
)

const BadContextErrorCode = 0
const QueryVersionOpcode = 0

type QueryVersionCookie uint64

const CreateContextOpcode = 1
const RegisterClientsOpcode = 2
const UnregisterClientsOpcode = 3
const GetContextOpcode = 4

type GetContextCookie uint64

const EnableContextOpcode = 5

type EnableContextCookie uint64

const DisableContextOpcode = 6
const FreeContextOpcode = 7

var readErrorFuncMap = make(map[uint8]x.ReadErrorFunc, 1)

func init() {
	readErrorFuncMap[BadContextErrorCode] = readBadContextError
	_ext = x.NewExtension("RECORD", 0, readErrorFuncMap)
}
