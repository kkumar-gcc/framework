package log

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	configmock "github.com/goravel/framework/contracts/config/mocks"
	"github.com/goravel/framework/support/carbon"
	"github.com/goravel/framework/support/file"
)

var singleLog = "storage/logs/goravel.log"
var dailyLog = fmt.Sprintf("storage/logs/goravel-%s.log", carbon.Now().ToDateString())

type LogrusTestSuite struct {
	suite.Suite
}

func TestLogrusTestSuite(t *testing.T) {
	suite.Run(t, new(LogrusTestSuite))
}

func (s *LogrusTestSuite) SetupTest() {

}

func (s *LogrusTestSuite) TestLogrus() {
	var (
		mockConfig *configmock.Config
		log        *Application
	)

	beforeEach := func() {
		mockConfig = initMockConfig()
	}

	tests := []struct {
		name   string
		setup  func()
		assert func()
	}{
		{
			name: "WithContext",
			setup: func() {
				mockConfig.On("GetString", "logging.channels.daily.level").Return("debug").Once()
				mockConfig.On("GetString", "logging.channels.single.level").Return("debug").Once()

				log = NewLogrusApplication(mockConfig)
			},
			assert: func() {
				writer := log.WithContext(context.Background())
				assert.Equal(s.T(), reflect.TypeOf(writer).String(), reflect.TypeOf(&Writer{}).String())
			},
		},
		{
			name: "Debug",
			setup: func() {
				mockDriverConfig(mockConfig)

				log = NewLogrusApplication(mockConfig)
				log.Debug("Debug Goravel")
			},
			assert: func() {
				assert.True(s.T(), file.Contain(singleLog, "test.debug: Debug Goravel"))
				assert.True(s.T(), file.Contain(dailyLog, "test.debug: Debug Goravel"))
			},
		},
		{
			name: "No Debug",
			setup: func() {
				mockConfig.On("GetString", "logging.channels.daily.level").Return("info").Once()
				mockConfig.On("GetString", "logging.channels.single.level").Return("info").Once()

				log = NewLogrusApplication(mockConfig)
				log.Debug("No Debug Goravel")
			},
			assert: func() {
				assert.False(s.T(), file.Contain(singleLog, "test.debug: No Debug Goravel"))
				assert.False(s.T(), file.Contain(dailyLog, "test.debug: No Debug Goravel"))
			},
		},
		{
			name: "Debugf",
			setup: func() {
				mockDriverConfig(mockConfig)

				log = NewLogrusApplication(mockConfig)
				log.Debugf("Goravel: %s", "World")
			},
			assert: func() {
				assert.True(s.T(), file.Contain(singleLog, "test.debug: Goravel: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.debug: Goravel: World"))
			},
		},
		{
			name: "Info",
			setup: func() {
				mockDriverConfig(mockConfig)

				log = NewLogrusApplication(mockConfig)
				log.Info("Goravel")
			},
			assert: func() {
				assert.True(s.T(), file.Contain(singleLog, "test.info: Goravel"))
				assert.True(s.T(), file.Contain(dailyLog, "test.info: Goravel"))
			},
		},
		{
			name: "Infof",
			setup: func() {
				mockDriverConfig(mockConfig)

				log = NewLogrusApplication(mockConfig)
				log.Infof("Goravel: %s", "World")
			},
			assert: func() {
				assert.True(s.T(), file.Contain(singleLog, "test.info: Goravel: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.info: Goravel: World"))
			},
		},
		{
			name: "Warning",
			setup: func() {
				mockDriverConfig(mockConfig)

				log = NewLogrusApplication(mockConfig)
				log.Warning("Goravel")
			},
			assert: func() {
				assert.True(s.T(), file.Contain(singleLog, "test.warning: Goravel"))
				assert.True(s.T(), file.Contain(dailyLog, "test.warning: Goravel"))
			},
		},
		{
			name: "Warningf",
			setup: func() {
				mockDriverConfig(mockConfig)

				log = NewLogrusApplication(mockConfig)
				log.Warningf("Goravel: %s", "World")
			},
			assert: func() {
				assert.True(s.T(), file.Contain(singleLog, "test.warning: Goravel: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.warning: Goravel: World"))
			},
		},
		{
			name: "Error",
			setup: func() {
				mockDriverConfig(mockConfig)

				log = NewLogrusApplication(mockConfig)
				log.Error("Goravel")
			},
			assert: func() {
				assert.True(s.T(), file.Contain(singleLog, "test.error: Goravel"))
				assert.True(s.T(), file.Contain(dailyLog, "test.error: Goravel"))
			},
		},
		{
			name: "Errorf",
			setup: func() {
				mockDriverConfig(mockConfig)

				log = NewLogrusApplication(mockConfig)
				log.Errorf("Goravel: %s", "World")
			},
			assert: func() {
				assert.True(s.T(), file.Contain(singleLog, "test.error: Goravel: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.error: Goravel: World"))
			},
		},
		{
			name: "Panic",
			setup: func() {
				mockDriverConfig(mockConfig)

				log = NewLogrusApplication(mockConfig)
			},
			assert: func() {
				assert.Panics(s.T(), func() {
					log.Panic("Goravel")
				})
				assert.True(s.T(), file.Contain(singleLog, "test.panic: Goravel"))
				assert.True(s.T(), file.Contain(dailyLog, "test.panic: Goravel"))
			},
		},
		{
			name: "Panicf",
			setup: func() {
				mockDriverConfig(mockConfig)

				log = NewLogrusApplication(mockConfig)
			},
			assert: func() {
				assert.Panics(s.T(), func() {
					log.Panicf("Goravel: %s", "World")
				})
				assert.True(s.T(), file.Contain(singleLog, "test.panic: Goravel: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.panic: Goravel: World"))
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			beforeEach()
			test.setup()
			test.assert()

			mockConfig.AssertExpectations(s.T())
		})
	}

	_ = file.Remove("storage")
}

func TestLogrus_Fatal(t *testing.T) {
	mockConfig := initMockConfig()
	mockDriverConfig(mockConfig)
	log := NewLogrusApplication(mockConfig)

	if os.Getenv("FATAL") == "1" {
		log.Fatal("Goravel")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLogrus_Fatal")
	cmd.Env = append(os.Environ(), "FATAL=1")
	err := cmd.Run()

	assert.EqualError(t, err, "exit status 1")
	assert.True(t, file.Contain(singleLog, "test.fatal: Goravel"))
	assert.True(t, file.Contain(dailyLog, "test.fatal: Goravel"))
	_ = file.Remove("storage")
}

func TestLogrus_Fatalf(t *testing.T) {
	mockConfig := initMockConfig()
	mockDriverConfig(mockConfig)
	log := NewLogrusApplication(mockConfig)

	if os.Getenv("FATAL") == "1" {
		log.Fatalf("Goravel")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLogrus_Fatal")
	cmd.Env = append(os.Environ(), "FATAL=1")
	err := cmd.Run()

	assert.EqualError(t, err, "exit status 1")
	assert.True(t, file.Contain(singleLog, "test.fatal: Goravel"))
	assert.True(t, file.Contain(dailyLog, "test.fatal: Goravel"))
	_ = file.Remove("storage")
}

func initMockConfig() *configmock.Config {
	mockConfig := &configmock.Config{}

	mockConfig.On("GetString", "logging.default").Return("stack").Once()
	mockConfig.On("GetString", "logging.channels.stack.driver").Return("stack").Once()
	mockConfig.On("Get", "logging.channels.stack.channels").Return([]string{"single", "daily"}).Once()
	mockConfig.On("GetString", "logging.channels.daily.driver").Return("daily").Once()
	mockConfig.On("GetString", "logging.channels.daily.path").Return(singleLog).Once()
	mockConfig.On("GetInt", "logging.channels.daily.days").Return(7).Once()
	mockConfig.On("GetBool", "logging.channels.daily.print").Return(false).Once()
	mockConfig.On("GetString", "logging.channels.single.driver").Return("single").Once()
	mockConfig.On("GetString", "logging.channels.single.path").Return(singleLog).Once()
	mockConfig.On("GetBool", "logging.channels.single.print").Return(false).Once()

	return mockConfig
}

func mockDriverConfig(mockConfig *configmock.Config) {
	mockConfig.On("GetString", "logging.channels.daily.level").Return("debug").Once()
	mockConfig.On("GetString", "logging.channels.single.level").Return("debug").Once()
	mockConfig.On("GetString", "app.timezone").Return("UTC")
	mockConfig.On("GetString", "app.env").Return("test")
}
