package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Project struct {
	Title          string
	Description    string
	TechStack      []string
	Image          string
	Link           string
	Screenshots    []string
	IsConfidential bool
	IsLocked       bool
}

type Experience struct {
	Year     string
	Role     string
	Company  string
	Location string
	Details  string
}

func Home(c *fiber.Ctx) error {
	experiences := []Experience{
		{
			Year:     "Jan 2024 - Present",
			Role:     "Systems Developer & Data Analyst (Intern)",
			Company:  "Ministry of Trade (Bappebti)",
			Location: "Jakarta, Indonesia",
			Details:  "Developing a transaction management system for commodity futures trading via Laravel 12 & Chart.js. conducting data analysis on legal vs. illegal trading entities to enhance regulatory insights.",
		},
		{
			Year:     "Sep 2023 - Dec 2023",
			Role:     "Freelance Systems Developer",
			Company:  "PT Kumaitu Kargo",
			Location: "Remote",
			Details:  "Built a digital pre-departure safety (K3) verification system for cargo drivers using Google Cloud (Sheets/Forms) as a centralized database, replacing manual paperwork.",
		},
		{
			Year:     "Jun 2023 - Aug 2023",
			Role:     "Java Developer Intern (KKP)",
			Company:  "PT Malik Hidayatullah",
			Location: "Jakarta, Indonesia",
			Details:  "Developed SIKARGO, a desktop application for managing warehouse cargo data, primarily focusing on recording shipping and receiving information.",
		},
	}

	projects := []Project{
		{
			Title:       "SIKARGO (Cargo Information System)",
			Description: "A desktop application used to manage and record daily warehouse cargo data, including tracking the status of stock for shipping and receiving operations.",
			TechStack:   []string{"Java", "Desktop App", "MySQL"},
			Image:       "/img/sikargo/halaman login.png",
			Link:        "#",
			Screenshots: []string{
				"/img/sikargo/halaman login.png",
				"/img/sikargo/dashboard.png",
				"/img/sikargo/tampilan data master.png",
				"/img/sikargo/tampilan data client.png",
				"/img/sikargo/tampilan data supir.png",
				"/img/sikargo/Tampilan data pengiriman.png",
				"/img/sikargo/tampilan data penerimaan.png",
				"/img/sikargo/Tampilan proses.png",
				"/img/sikargo/tampilan stok barang.png",
				"/img/sikargo/Tampilan report.png",
			},
		},
		{
			Title:          "Bappebti Legality Check Log Analysis",
			Description:    "Analyzed the Bappebti legality check web portal's search logs. Cleansed raw data using Python, and designed an interactive dashboard in Looker Studio to track web performance and identify user search trends for both legal and illegal trading entities.",
			TechStack:      []string{"Python", "Data Cleansing", "Looker Studio"},
			Image:          "/img/analisis data  log pencarian/executive summary.png",
			Link:           "#",
			Screenshots:    []string{"/img/analisis data  log pencarian/executive summary.png"},
			IsConfidential: true,
		},
		{
			Title:       "SIDATA (Exchange Data Management)",
			Description: "A web platform named SIDATA developed to manage and visualize monthly and annual transaction reports from commodity exchanges. Built with Laravel 12 and Bootstrap, featuring interactive data visualizations powered by Chart.js.",
			TechStack:   []string{"Laravel 12", "Bootstrap", "Chart.js", "MySQL"},
			Image:       "/img/sidataapp/sidata.png",
			Link:        "#",
			Screenshots: []string{"/img/sidataapp/sidata.png", "/img/sidataapp/sidata.png"},
			IsLocked:    true,
		},
	}

	log.Printf("Rendering Home with %d projects and %d experiences", len(projects), len(experiences))

	return c.Render("home", fiber.Map{
		"Title":       "Agung Prayogi | Full-Stack Developer & Data Analyst",
		"Projects":    projects,
		"Experiences": experiences,
	}, "layouts/main")
}
