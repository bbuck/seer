package config

// Unforunately there is no way to package static binary files in Go, so the
// next best thing to provide a "default" config file is to simply do it
// this way where we define a string and use that.
const defaultConfigToml = `# Config version 1.0

# If you're hosting more than one server out of the graph database this field
# should be set to the ID of the root node of this application. Otherwise the
# start node will be fetched with the following Cypher:
#
#    MATCH (s:Start) RETURN s LIMIT 1
#
#root_node_id = ""

# Define the information necessary to connect to the Datatabase. At this time
# Authentication is not supported so this is simply the host of the web server
# which, for security, should be on an internal only machine.
[database]

# URL to the Neo4j instance. The default address is "localhost:7474" so this
# can remain commented unless the path is different.
#host = "localhost:7474"
`
