package upsilonAdaptorMermaid;

func BuildMermaidGraph() string {
	ret := "graph LR\n";

	ret += "AMQP\n";
	ret += "DB(fa:fa-database DB)\n";
	ret += "DB --> web\n"
	ret += "web --> DB\n"
	ret += "web --> AMQP\n"

	//ret += "classDef drone\n"
	ret += "classDef karmaGood fill:#90ee90\n"
	ret += "classDef karmaBad fill:salmon\n"

	for _, node := range GetNodes("all") {
		ret += node.Identifier + "\n"

		switch karma := node.Karma; karma {
		case "GOOD":
			ret += "class " + node.Identifier + " karmaGood\n";
		case "BAD":
			ret += "class " + node.Identifier + " karmaBad\n";
		}
	}

	for _, drone := range GetNodes("drone") {
		ret += "class " + drone.Identifier + " drone\n"
		ret += drone.Identifier + " --> " + "AMQP\n"
	}

	for _, cust := range GetNodes("custodian") {
		ret += "AMQP --> " + cust.Identifier + "\n";
		ret += cust.Identifier + " --> DB\n"
	}

	for _, reactor := range GetNodes("reactor") {
		ret += "AMQP --> " + reactor.Identifier + "\n";
	}

	return ret;
}
