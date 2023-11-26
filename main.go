package main

import (
	"encoding/json"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/xsen84/goLibreFollower/config"
	"github.com/xsen84/goLibreFollower/utils"
)

var settingsWindowOn bool
var a fyne.App

func main() {

	// define path to configuration file
	homeDir, _ := os.UserHomeDir()
	configurationFile := filepath.Join(homeDir, "glf-settings.json")

	var lcfg = config.AppConfig{}
	lcfg = readSettings(a, configurationFile)

	a := app.New()
	w := a.NewWindow("goLibreFollower " + config.Version)
	a.Settings().SetTheme(theme.DarkTheme())
	a.SetIcon(config.AppIcon)

	glucosePanel := canvas.NewText("0", config.DefaultColor)
	glucosePanel.TextSize = lcfg.TextSize

	diffPanel := canvas.NewText("0", config.DefaultColor)
	diffPanel.TextSize = lcfg.TextSize / 3

	//trendPanel := widget.NewLabel("")

	lastReadingPanel := canvas.NewText("0", config.DefaultColor)
	lastReadingPanel.TextSize = lcfg.TextSize / 3

	statusPanel := canvas.NewText("authenticating", config.DefaultColor)
	statusPanel.TextSize = lcfg.TextSize / 7

	//read png end show as bytes
	// x, _ := os.ReadFile("appicon.ico")
	// fmt.Println("-----")
	// fmt.Println([]byte(x))
	// fmt.Println("-----")

	settingsWindowOn = false
	buttonSettings := widget.NewButtonWithIcon("", config.SettingsButtonIcon, func() {
		if !settingsWindowOn {
			settingsWindowOn = true
			setSettings(configurationFile, a)
		}
	})

	content2 := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), diffPanel, lastReadingPanel, layout.NewSpacer())
	content1 := container.New(layout.NewHBoxLayout(), glucosePanel, layout.NewSpacer(), content2, layout.NewSpacer())
	content3 := container.New(layout.NewVBoxLayout(), content1, layout.NewSpacer(), statusPanel)
	content4 := container.New(layout.NewHBoxLayout(), content3, buttonSettings)

	w.SetContent(content4)

	// authenticating
	isAuthSuccessful, authToken := getAuthToken(lcfg.Username, lcfg.Password, lcfg.Region)
	if isAuthSuccessful {
		statusPanel.Text = "Authentication successfull"
		go func() {
			connectAndGetData(authToken, lcfg.RefreshInterval, glucosePanel, diffPanel, lastReadingPanel, statusPanel)
		}()
	} else {
		statusPanel.Text = authToken
	}
	w.Show()
	a.Run()
}

func readSettings(a fyne.App, configurationFile string) config.AppConfig {
	var lcfg config.AppConfig
	settings, err := ioutil.ReadFile(configurationFile)
	if err != nil {
		log.Print("error opening file")
		setSettings(configurationFile, a)
	}
	err = json.Unmarshal(settings, &lcfg)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return lcfg
}

func setSettings(configurationFile string, a fyne.App) {
	swindow := a.NewWindow("Settings")
	swindow.SetOnClosed(func() {
		settingsWindowOn = false
		swindow.Close()
	})
	nameText := canvas.NewText(" Setting                                                   ", config.DefaultColor)
	descText := canvas.NewText(" Description", config.DefaultColor)

	usernameLabel := widget.NewLabel("Follower email:")
	usernameInput := widget.NewEntry()
	usernameInput.SetPlaceHolder("Username")

	passwordLabel := widget.NewLabel("Follower Password: ")
	passwordInput := widget.NewEntry()
	passwordInput.SetPlaceHolder("Password")

	regionLabel := widget.NewLabel("LibreLinkUp Region:")
	regionInput := widget.NewEntry()
	regionInput.SetPlaceHolder("Region - default: de")

	textSizeLabel := widget.NewLabel("Text Size:")
	textSizeInput := widget.NewEntry()
	textSizeInput.SetPlaceHolder("Text Size")

	refreshIntervalLabel := widget.NewLabel("Refresh Interval:")
	refreshIntervalInput := widget.NewEntry()
	refreshIntervalInput.SetPlaceHolder("Refresh Interfal - deafult: 10s")

	buttonSave := widget.NewButton("Save", func() {
		var lcfg = config.AppConfig{}
		lcfg.Username = usernameInput.Text
		lcfg.Password = passwordInput.Text
		if regionInput.Text == "" {
			lcfg.Region = "de"
		} else {
			lcfg.Region = regionInput.Text
		}
		if textSizeInput.Text == "" {
			lcfg.TextSize = 100
		} else {
			ts, _ := strconv.Atoi(textSizeInput.Text)
			lcfg.TextSize = float32(ts)
		}
		if refreshIntervalInput.Text == "" {
			lcfg.RefreshInterval = 10
		} else {
			lcfg.RefreshInterval, _ = strconv.Atoi(refreshIntervalInput.Text)
		}
		jsonSettings, err := json.Marshal(lcfg)
		if err != nil {
			log.Fatal("Error during Marshal(): ", err)
		}
		err = ioutil.WriteFile(configurationFile, jsonSettings, 0644)
		if err != nil {
			log.Fatal("Error during Save: ", err)
		}
		settingsWindowOn = false
		swindow.Close()
	})

	// try to load config
	var lcfg = config.AppConfig{}

	settings, err := ioutil.ReadFile(configurationFile)
	if err != nil {
		log.Print("No configuration found")
	} else {
		// try to read values
		err = json.Unmarshal(settings, &lcfg)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
		usernameInput.Text = lcfg.Username
		passwordInput.Text = lcfg.Password
		regionInput.Text = lcfg.Region
		textSizeInput.Text = strconv.Itoa(int(lcfg.TextSize))
		refreshIntervalInput.Text = strconv.Itoa(lcfg.RefreshInterval)

	}

	centeredButtonSave := container.New(layout.NewCenterLayout(), buttonSave)
	infoText := canvas.NewText("After SAVE please restart application !!!", color.RGBA{255, 0, 0, 255})
	centeredInfoText := container.New(layout.NewCenterLayout(), infoText)

	contentLabels := container.New(layout.NewVBoxLayout(), descText, usernameLabel, passwordLabel, regionLabel, textSizeLabel, refreshIntervalLabel)
	contentInputs := container.New(layout.NewVBoxLayout(), nameText, usernameInput, passwordInput, regionInput, textSizeInput, refreshIntervalInput)
	settingsContent := container.New(layout.NewHBoxLayout(), contentLabels, contentInputs)
	allContents := container.New(layout.NewVBoxLayout(), settingsContent, centeredButtonSave, centeredInfoText)
	swindow.SetContent(allContents)
	swindow.Show()
	a.Run()

}

func connectAndGetData(authToken string, refreshInterval int, glucosePanel *canvas.Text, diffPanel *canvas.Text, lastReadingPanel *canvas.Text, statusPanel *canvas.Text) {
	oldGlucose := 0
	oldTimeStamp := ""
	lastDiff := 0
	intervals := 0

	for i := 1; i > 0; i++ {
		output := ""
		httpError := false
		statusPanel.Color = config.DefaultColor
		glucosePanel.Color = config.BlueColor
		connResponseResult, err := utils.GetReadings(authToken)
		if err != nil {
			statusPanel.Text = err.Error()
			httpError = true
		}

		connData := config.ConnResponse{}
		json.Unmarshal([]byte(connResponseResult), &connData)
		if connData.Status != 0 {
			statusPanel.Text = connData.Error.Message
		}

		if (connData.Status == 0) && !httpError && (len(connData.Data) != 0) {
			if connData.Data[0].GlucoseMeasurement.Timestamp != oldTimeStamp {
				lastDiff = connData.Data[0].GlucoseMeasurement.Value - oldGlucose
				oldTimeStamp = connData.Data[0].GlucoseMeasurement.Timestamp
				intervals = 0
			} else {
				intervals++
			}
			statusPanel.Text = "Last value recorded at: " + connData.Data[0].GlucoseMeasurement.Timestamp
			if lastDiff > 0 {
				diffPanel.Text = "+" + strconv.Itoa(lastDiff)
			} else {
				diffPanel.Text = strconv.Itoa(lastDiff)
			}
			date, _ := time.Parse("1/2/2006 15:04:05 PM", connData.Data[0].GlucoseMeasurement.Timestamp)

			// add the TZ difference
			date = date.Add(-2 * time.Hour)
			timeSince := utils.HumanizeDuration(time.Since(date))
			output += strconv.Itoa(connData.Data[0].GlucoseMeasurement.Value)
			lastReadingPanel.Text = timeSince
			oldGlucose = connData.Data[0].GlucoseMeasurement.Value
			glucosePanel.Text = output
			if connData.Data[0].GlucoseMeasurement.Value < config.LowGlucose {
				glucosePanel.Color = config.RedColor
			}
			if connData.Data[0].GlucoseMeasurement.Value > config.HighGlucose {
				glucosePanel.Color = config.YellowColor
			}

			if (connData.Data[0].GlucoseMeasurement.Value > config.LowGlucose) && (connData.Data[0].GlucoseMeasurement.Value < config.HighGlucose) {
				glucosePanel.Color = config.GreenColor
			}
			if intervals > config.TimeoutInterval {
				statusPanel.Text = "Check sensor for errors or phone conectivity"
				statusPanel.Color = config.RedColor
			}
		} else {
			statusPanel.Text = "Failded to load connections"
			statusPanel.Color = config.RedColor
		}

		glucosePanel.Refresh()
		diffPanel.Refresh()
		lastReadingPanel.Refresh()
		statusPanel.Refresh()
		time.Sleep(time.Duration(refreshInterval) * time.Second)
	}

}

func getAuthToken(username string, password string, region string) (bool, string) {
	authResponseResult, err := utils.Auth(username, password, region)
	if err != nil {
		return false, err.Error()
	}

	authData := config.AuthResponse{}

	json.Unmarshal([]byte(authResponseResult), &authData)
	if authData.Status != 0 {
		return false, authData.Error.Message
	}
	return true, authData.Data.AuthTicket.Token
}
