package main;

func BuildMermaidGraph() string {
	ret := "graph LR\n";

	ret += "AMQP\n";
	ret += "DB(fa:fa-database DB)\n";
	ret += "style DB fill:#dee3e7\n"
	ret += "DB --> web\n"
	ret += "web --> DB\n"
	ret += "web --> AMQP\n"

	ret += "classDef drone fill:#42f477\n"

	for _, node := range GetNodes("all") {
		ret += node + "\n"
	}

	for _, drone := range GetNodes("drone") {
		ret += "class " + drone + " drone\n"
		ret += drone + " --> " + "AMQP\n"
	}

	for _, cust := range GetNodes("custodian") {
		ret += "AMQP --> " + cust + "\n";
		ret += cust + " --> DB\n"
	}

	for _, reactor := range GetNodes("reactor") {
		ret += "AMQP --> " + reactor + "\n";
	}

	return ret;
}
