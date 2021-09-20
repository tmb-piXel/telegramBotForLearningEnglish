package config

import (
	"github.com/spf13/viper"
)

type Messages struct {
	Responses
}

type Responses struct {
	StartMessage        string `mapstructure:"start_message"`
	HelpMessage         string `mapstructure:"help_message"`
	AlreadyStart        string `mapstructure:"already_start"`
	UnknownCommand      string `mapstructure:"unknown_command"`
	CorrectAnswer       string `mapstructure:"correct_answer"`
	WrongAnswer         string `mapstructure:"wrong_answer"`
	TheCorrectAnswerWas string `mapstructure:"the_correct_answer_was"`
	SelectLanguage      string `mapstructure:"select_language"`
}

type Config struct {
	TelegramToken   string
	PathDictonaries string `mapstructure:"path_dictionaries"`
	PostgresqlUrl   string
	Messages        Messages
	Buttons         Buttons
}

type Buttons struct {
	SetLang  string `mapstructure:"set_lang"`
	Settings string `mapstructure:"settings"`
	Help     string `mapstructure:"help"`
	SetTopic string `mapstructure:"set_topic"`
	List     string `mapstructure:"list"`
	FromRu   string `mapstructure:"fromRu"`
	ToRu     string `mapstructure:"toRu"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("buttons", &cfg.Buttons); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	if err := viper.BindEnv("telegramtoken", "postgresqlurl"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("telegramtoken")
	cfg.TelegramToken = `1955006608:AAFfX0ca0VBLFs6uwVr9VdLZVHo3t1ikGj0`
	cfg.PostgresqlUrl = viper.GetString("postgresqlurl")
	cfg.PostgresqlUrl = `host=localhost port=5433 user=postgres password=postgres dbname=postgres sslmode=disable`

	return nil
}
