package kafka

type metadataResponse struct {
	brokers []Broker
	topics  []topicMetadata
}

func (m *metadataResponse) encode(pe packetEncoder) {
	pe.putInt32(int32(len(m.brokers)))
	for i := range m.brokers {
		(&m.brokers[i]).encode(pe)
	}
	pe.putInt32(int32(len(m.topics)))
	for i := range m.topics {
		(&m.topics[i]).encode(pe)
	}
}

func (m *metadataResponse) decode(pd packetDecoder) (err error) {
	n, err := pd.getArrayCount()
	if err != nil {
		return err
	}

	m.brokers = make([]Broker, n)
	for i := 0; i < n; i++ {
		err = (&m.brokers[i]).decode(pd)
		if err != nil {
			return err
		}
	}

	n, err = pd.getArrayCount()
	if err != nil {
		return err
	}

	m.topics = make([]topicMetadata, n)
	for i := 0; i < n; i++ {
		err = (&m.topics[i]).decode(pd)
		if err != nil {
			return err
		}
	}

	return nil
}
