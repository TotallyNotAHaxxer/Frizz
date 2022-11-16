package Engine

import (
	"encoding/json"
)

func Read__Store(Channel <-chan string) {
	// have this function load the json databases and continue
	for DatainChan := range Channel {
		switch DatainChan {
		case "Modules/Gsrc/Database/FrizzDatabase.json":
			// Load the marshall the Frizz structure
			ToUnMarhshalFrizz := ReturnJsonInterface(DatainChan)
			json.Unmarshal([]byte(ToUnMarhshalFrizz), &StructureFrizzPointer)
		case "Modules/Gsrc/Database/PreProcessor.json":
			// Load the preprocessor structure
			ToUnMarshalPreProcessor := ReturnJsonInterface(DatainChan)
			json.Unmarshal([]byte(ToUnMarshalPreProcessor), &StructurePreProcessor)
		case "Modules/Gsrc/Database/PCPP.json":
			// Load the pcpp structure
			ToUnMarshalPCPP := ReturnJsonInterface(DatainChan)
			json.Unmarshal([]byte(ToUnMarshalPCPP), &StructurePcapPlusPlus)
		}
	}
}

func Find__Send(Channel chan<- string) {
	wg.Add(1)
	go func() {
		for k := range Database {
			if ExistQ(Database[k]) {
				Channel <- Database[k]
			} else {
				//fmt.Println("[Engine] Error: Could not find database file | Data not sent to channel; CHUNK => ", Database[k])
			}
		}
		wg.Done()
	}()
	wg.Wait()
	close(Channel)
}
