package models

// ApplicationParams stores the global information associated with an application.
type ApplicationParams struct {
	// ApprovalProgram approval program.
	ApprovalProgram []byte `json:"approval-program"`

	// ClearStateProgram clear state program.
	ClearStateProgram []byte `json:"clear-state-program"`

	// Creator the address that created this application. This is the address where the
	// parameters and global state for this application can be found.
	Creator string `json:"creator,omitempty"`

	// ExtraProgramPages the number of extra program pages available to this app.
	ExtraProgramPages uint64 `json:"extra-program-pages,omitempty"`

	// GlobalState global state
	GlobalState []TealKeyValue `json:"global-state,omitempty"`

	// GlobalStateSchema global schema
	GlobalStateSchema ApplicationStateSchema `json:"global-state-schema,omitempty"`

	// LocalStateSchema local schema
	LocalStateSchema ApplicationStateSchema `json:"local-state-schema,omitempty"`

	// Version the number of updates to the application programs
	Version uint64 `json:"version,omitempty"`
}
