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
		case "Modules/Gsrc/Database/ApplicationInformationStorage.json":
			// Load the application info structure
			ToUnmarshalServerInfo := ReturnJsonInterface(DatainChan)
			json.Unmarshal([]byte(ToUnmarshalServerInfo), &StructureAppInfo)
		case "Modules/Gsrc/Database/ServerInfo.json":
			// Load server information
			ToUnmarshalServerInformation := ReturnJsonInterface(DatainChan)
			json.Unmarshal([]byte(ToUnmarshalServerInformation), &StructureServerInfo)
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
				//("[Engine] Error: Could not find database file | Data not sent to channel; CHUNK => ", Database[k])
			}
		}
		wg.Done()
	}()
	wg.Wait()
	close(Channel)
}
