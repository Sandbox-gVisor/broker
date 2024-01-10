package storages

const uri = "<connection string>"

type MongodbStorage struct {
}

func (store *MongodbStorage) Init() {
	//serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	//opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	//_, _ := mongo.Connect(context.TODO(), opts)
}
