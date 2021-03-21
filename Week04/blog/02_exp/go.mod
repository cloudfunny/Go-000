module example

go 1.15

require github.com/Sirupsen/logrus v1.4.1 // indirect

replace (
	github.com/Sirupsen/logrus v1.4.1 => github.com/sirupsen/logrus v1.8.0
)
