package utils

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/umer-emumba/BudgetBuddy/config"
	"github.com/umer-emumba/BudgetBuddy/repositories"
	"github.com/umer-emumba/BudgetBuddy/types"
)

func StartCronJob() {
	c := cron.New()
	// Schedule the cron job to run first day of month at 12am
	c.AddFunc("@monthly", MonthlyScheduler)

	// Start the cron job scheduler
	c.Start()

}

func MonthlyScheduler() {
	fmt.Printf("cron started")
	config.InitLogger()
	tranRepo := repositories.NewTransactionRepository()
	userRepo := repositories.NewUserRepository()
	users, userErr := userRepo.GetAllUsers()
	if userErr != nil {
		config.Logger.Error(userErr.Error())
	}

	for _, user := range users {
		report, reportErr := tranRepo.MonthlyReport(user.ID)
		if reportErr != nil {
			config.Logger.Error(reportErr.Error())
		} else {
			var mailBody string
			if len(report) > 0 {
				mailBody = "Your last month entries are following:<br><ul>"

				for _, item := range report {
					mailBody += fmt.Sprintf("<li> <strong>%s</strong> is %v </li>", item.TransactionType, item.TotalAmount)
				}
				mailBody += "</ul>"
			} else {
				mailBody = "You have done no transaction in last month"
			}

			mailOptions := types.MailOptions{
				To:      user.Email,
				Subject: "Monthly Report",
				Body:    mailBody,
			}

			client := config.CreateAsynqClient()
			defer client.Close()
			task, emailErr := NewEmailDeliveryTask(mailOptions)
			if emailErr != nil {
				config.Logger.Error("could not create task: " + emailErr.Error())
			}
			info, err := client.Enqueue(task)
			if err != nil {
				config.Logger.Error("could not enqueue task: " + err.Error())
			}
			config.Logger.Info("enqueued task: id=" + info.ID + " queue=" + info.Queue)
		}
	}
}
