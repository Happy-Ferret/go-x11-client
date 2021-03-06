package dpms

import "github.com/linuxdeepin/go-x11-client"

// #WREQ
func writeGetVersion(w *x.Writer, clientMajorVersion, clientMinorVersion uint16) {
	w.WritePad(4)
	w.Write2b(clientMajorVersion)
	w.Write2b(clientMinorVersion)
}

type GetVersionReply struct {
	ServerMajorVersion uint16
	ServerMinorVersion uint16
}

func readGetVersionReply(r *x.Reader, v *GetVersionReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ServerMajorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ServerMinorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeCapable(w *x.Writer) {
	w.WritePad(4)
}

type CapableReply struct {
	Capable bool
}

func readCapableReply(r *x.Reader, v *CapableReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Capable = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(23)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeGetTimeouts(w *x.Writer) {
	w.WritePad(4)
}

type GetTimeoutsReply struct {
	StandbyTimeout uint16
	SuspendTimeout uint16
	OffTimeout     uint16
}

func readGetTimeoutsReply(r *x.Reader, v *GetTimeoutsReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.StandbyTimeout = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SuspendTimeout = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.OffTimeout = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(18)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeSetTimeouts(w *x.Writer, standbyTimeout, suspendTimeout, offTimeout uint16) {
	w.WritePad(4)
	w.Write2b(standbyTimeout)
	w.Write2b(suspendTimeout)
	w.Write2b(offTimeout)
	w.WritePad(2)
}

// #WREQ
func writeEnable(w *x.Writer) {
	w.WritePad(4)
}

// #WREQ
func writeDisable(w *x.Writer) {
	w.WritePad(4)
}

// #WREQ
func writeForceLevel(w *x.Writer, powerLevel uint16) {
	w.WritePad(4)
	w.Write2b(powerLevel)
	w.WritePad(2)
}

// #WREQ
func writeInfo(w *x.Writer) {
	w.WritePad(4)
}

type InfoReply struct {
	PowerLevel uint16
	State      bool
}

func readInfoReply(r *x.Reader, v *InfoReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.PowerLevel = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.State = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(21)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}
