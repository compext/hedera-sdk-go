package hedera

type AddressBookQuery struct {
	attempt     uint64
	maxAttempts uint64
	fileID      *FileID
	limit       int32
}

func NewAddressBookQuery() *AddressBookQuery {
	return &AddressBookQuery{
		fileID: nil,
		limit:  0,
	}
}

func (query *AddressBookQuery) SetFileID(id FileID) *AddressBookQuery {
	query.fileID = &id
	return query
}

func (query *AddressBookQuery) GetFileID() FileID {
	if query.fileID == nil {
		return FileID{}
	}

	return *query.fileID
}

func (query *AddressBookQuery) SetLimit(limit int32) *AddressBookQuery {
	query.limit = limit
	return query
}

func (query *AddressBookQuery) GetLimit() int32 {
	return query.limit
}

func (query *AddressBookQuery) SetMaxAttempts(maxAttempts uint64) *AddressBookQuery {
	query.maxAttempts = maxAttempts
	return query
}

func (query *AddressBookQuery) GetMaxAttempts() uint64 {
	return query.maxAttempts
}

func (query *AddressBookQuery) _ValidateNetworkOnIDs(client *Client) error {
	if client == nil || !client.autoValidateChecksums {
		return nil
	}

	if query.fileID != nil {
		if err := query.fileID.ValidateChecksum(client); err != nil {
			return err
		}
	}

	return nil
}

func (query *AddressBookQuery) Execute(client *Client) (NodeAddressBook, error) {
	err := query._ValidateNetworkOnIDs(client)
	if err != nil {
		return NodeAddressBook{}, err
	}

	return NodeAddressBook{
		NodeAddresses: []NodeAddress{},
	}, nil
}
