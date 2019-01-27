# Sample bulk loader for Horde

This program is a sample bulk loader for adding devices to a given
collection by reading them from a CSV file.  The CSV file must
contain, in order, the following fields:

  - Device name
  - IMEI
  - IMSI
  
A sample input file might look something like this:

    "Test Device 1",111122223333,555566667777
    "Test Device 2",222233334444,666677778888

# Building

This program depends on the NB-IoT client clibrary for the NB-IoT
service, which can be installed by issuing the command:

    go get -u github.com/telenordigital/nbiot-go
	
You can then build the program by running

    go build
	
# Configuring access

Before running this program you must generate an appropriate API Token
and either put the token into the `.telenor-nbiot` file or set the
`TELENOR_NBIOT_TOKEN` environment variable.

See the documentation for
the [Go Library](https://github.com/telenordigital/nbiot-go) for more
information about the configuration file.

You can create an API Token by going to the "API tokens" menu on
https://nbiot.engineering/ and following the instructions.  It is
recommended that you limit the scope of your API Tokens so they have
as limited privileges as possible.


# Running

Once you have created an API Token and either created a config file or
set the environment variable you find the Collection ID you want to
load devices into via the user interface and then run the binary with
the following command line options:

    ./horde-device-loader -csv <CSV filename> -collection <collection id>
	

# Caveats

This is, as mentioned above, an example of how you can write a bulk
loader just to demonstrate how the API can be used.  In a real
bulk-loader you want to take a bit more care with the output so that
it is easier to write a parser for the output.
