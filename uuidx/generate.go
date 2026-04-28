package uuidx

import "github.com/google/uuid"

// Namespace definitions for V3 and V5
var (
	NamespaceDNS  = uuid.NameSpaceDNS.String()
	NamespaceURL  = uuid.NameSpaceURL.String()
	NamespaceOID  = uuid.NameSpaceOID.String()
	NamespaceX500 = uuid.NameSpaceX500.String()
)

// V1 generates a version 1 UUID (time-based)
func V1() string {
	u, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil.String()
	}

	return u.String()
}

// V2 generates a version 2 UUID (DCE Security - Person)
func V2() string {
	u, err := uuid.NewDCEPerson()
	if err != nil {
		return uuid.Nil.String()
	}

	return u.String()
}

// V3 generates a version 3 UUID (name-based MD5)
func V3(namespace string, name string) string {
	ns, err := uuid.Parse(namespace)
	if err != nil {
		return uuid.Nil.String()
	}

	return uuid.NewMD5(ns, []byte(name)).String()
}

// V4 generates a version 4 UUID (random)
func V4() string {
	return uuid.NewString()
}

// V5 generates a version 5 UUID (name-based SHA-1)
func V5(namespace string, name string) string {
	ns, err := uuid.Parse(namespace)
	if err != nil {
		return uuid.Nil.String()
	}

	return uuid.NewSHA1(ns, []byte(name)).String()
}

// V6 generates a version 6 UUID (reordered time-based)
func V6() string {
	u, err := uuid.NewV6()
	if err != nil {
		return uuid.Nil.String()
	}

	return u.String()
}

// V7 generates a version 7 UUID (Unix Epoch time-based)
func V7() string {
	u, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil.String()
	}

	return u.String()
}
