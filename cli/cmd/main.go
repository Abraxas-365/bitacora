package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWidth() uint {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return uint(ws.Col)
}

func main() {

	// Agregar flags al comando raíz
	tags := []string{}
	title := ""
	description := ""
	nickname := ""
	tagsString := ""
	error := ""
	solution := ""
	from := 0
	size := 0
	id := ""

	token := os.Getenv("BITACORA")
	url := os.Getenv("BITACORAURL")
	if url == "" {
		url = "http://localhost:1234"
	}

	type Report struct {
		Id          *string   `json:"id"`
		Tags        *[]string `json:"tags"`
		Title       *string   `json:"title"`
		Description *string   `json:"description"`
		Error       *string   `json:"error"`
		Solution    *string   `json:"solution"`
		Status      *bool     `json:"status"`
		Nickname    *string   `json:"nickname"`
	}

	type ReportQuery struct {
		Tags        *string `json:"tags"`
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Error       *string `json:"error"`
		Status      *bool   `json:"status"`
		Nickname    *string `json:"nickname"`
		From        int     `json:"from"`
		Size        int     `json:"size"`
	}

	// Inicializar el comando raíz
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "Una aplicación de ejemplo CLI",
		Long:  "Esta es una aplicación de ejemplo CLI que demuestra cómo usar la biblioteca Cobra en Go",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	var postCmd = &cobra.Command{
		Use:   "post",
		Short: "Una aplicación de ejemplo CLI",
		Long:  "Esta es una aplicación de ejemplo CLI que demuestra cómo usar la biblioteca Cobra en Go",

		Run: func(cmd *cobra.Command, args []string) {

			reportPost := Report{}
			if tags != nil {
				reportPost.Tags = &tags
			}
			if title != "" {
				reportPost.Title = &title
			}
			if description != "" {
				reportPost.Description = &description
			}
			if error != "" {
				reportPost.Error = &error
			}
			if solution != "" {
				reportPost.Solution = &solution
			}

			var client http.Client
			var buf bytes.Buffer

			if err := json.NewEncoder(&buf).Encode(reportPost); err != nil {
				log.Fatal(err)
			}

			req, err := http.NewRequest("POST", url+"/report", &buf)
			if err != nil {
				log.Fatal(err)
			}
			req.Header.Add("Authorization", "Bearer "+token)
			req.Header.Add("Content-Type", "application/json")
			resp, err := client.Do(req)
			if resp.StatusCode == 201 {
				color.Green("Se guardo con exito")
			} else if resp.StatusCode == 500 {
				color.Red("Hubo un Problema")
			}
			if err != nil {
				log.Fatal(err)
			}

		},
	}

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Una aplicación de ejemplo CLI",
		Long:  "Esta es una aplicación de ejemplo CLI que demuestra cómo usar la biblioteca Cobra en Go",

		Run: func(cmd *cobra.Command, args []string) {

			reportQuery := ReportQuery{}
			if tagsString != "" {
				reportQuery.Tags = &tagsString
			}
			if title != "" {
				reportQuery.Title = &title
			}
			if description != "" {
				reportQuery.Description = &description
			}
			if error != "" {
				reportQuery.Error = &error
			}
			if nickname != "" {
				reportQuery.Nickname = &nickname
			}
			reportQuery.From = from
			reportQuery.Size = size

			var client http.Client
			var buf bytes.Buffer
			var respRepotrs []Report

			if err := json.NewEncoder(&buf).Encode(reportQuery); err != nil {
				log.Println(err.Error())
			}

			fmt.Println()
			req, err := http.NewRequest("GET", url+"/report", &buf)
			if err != nil {
				log.Println(err.Error())
			}
			req.Header.Add("Authorization", "Bearer "+token)
			req.Header.Add("Content-Type", "application/json")
			resp, err := client.Do(req)

			if resp.StatusCode == 500 {
				color.Red("Hubo un Problema")
				return

			}
			if err != nil {
				log.Println(err.Error())
			}
			body, err := ioutil.ReadAll(resp.Body)

			if err := json.Unmarshal(body, &respRepotrs); err != nil {
				log.Println(err.Error())
			}
			winSize := getWidth()
			for _, report := range respRepotrs {

				fmt.Print("\n")
				for x := 0; x < int(winSize); x++ {
					fmt.Print("-")
				}
				fmt.Print("\n")
				if *report.Id != "" {
					fmt.Print(*report.Id)
				}
				fmt.Print("\n")
				if *report.Title != "" {
					fmt.Print(color.HiMagentaString("TITULO: "), *report.Title, "  ")
				}
				if report.Tags != nil {
					fmt.Print(color.CyanString("Tags: "), color.CyanString(strings.Join(*report.Tags, ",")), "  \n")
				}
				if *report.Description != "" {
					fmt.Print("\n")
					color.Green("Description: ")
					color.Blue("" + *report.Description)
				}
				if *report.Error != "" {
					fmt.Print("\n")
					color.Green("Error: ")
					color.Red("" + *report.Error)
				}
				if *report.Solution != "" {
					fmt.Print("\n")
					color.Green("Solucion: ")
					color.Blue("" + *report.Solution)
				}
				if *report.Nickname != "" {
					fmt.Print("\n")
					color.Green("Author: " + *report.Nickname)
				}
			}

			fmt.Print("\n")
			for x := 0; x < int(winSize); x++ {
				fmt.Print("-")

			}

		},
	}

	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a report",
		Long:  "Update a report",
		Run: func(cmd *cobra.Command, args []string) {

			reportPost := Report{}
			if len(tags) > 0 {
				reportPost.Tags = &tags
			}
			if title != "" {
				reportPost.Title = &title
			}
			if description != "" {
				reportPost.Description = &description
			}
			if error != "" {
				reportPost.Error = &error
			}
			if solution != "" {
				reportPost.Solution = &solution
			}

			var client http.Client
			var buf bytes.Buffer

			if err := json.NewEncoder(&buf).Encode(reportPost); err != nil {
				log.Println(err.Error())
			}

			req, err := http.NewRequest("PUT", url+"/report/"+id, &buf)
			if err != nil {
				log.Println(err.Error())
			}
			req.Header.Add("Authorization", "Bearer "+token)
			req.Header.Add("Content-Type", "application/json")
			resp, err := client.Do(req)
			if resp.StatusCode == 200 {
				color.Green("Se actualizo con exito")
			} else if resp.StatusCode == 500 {
				color.Red("Hubo un Problema")
			}
			if err != nil {
				log.Println(err.Error())
			}
		},
	}
	// Ejecutar la aplicación

	postCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "Tags separados por comas")
	postCmd.Flags().StringVarP(&title, "title", "T", "", "Título post")
	postCmd.Flags().StringVarP(&error, "error", "e", "", "Descripción del error")
	postCmd.Flags().StringVarP(&solution, "solution", "s", "", "Solucion del error")
	postCmd.Flags().StringVarP(&description, "description", "d", "", "Descripción")
	rootCmd.AddCommand(postCmd)

	getCmd.Flags().StringVarP(&tagsString, "tags", "t", "", "Tags separados por comas")
	getCmd.Flags().StringVarP(&title, "title", "T", "", "Título del post")
	getCmd.Flags().StringVarP(&error, "error", "e", "", "Descripción del error")
	getCmd.Flags().StringVarP(&description, "description", "d", "", "Descripción")
	getCmd.Flags().StringVarP(&nickname, "nickname", "n", "", "Apodo del creador")
	getCmd.Flags().IntVarP(&from, "from", "f", 0, "Desde")
	getCmd.Flags().IntVarP(&size, "size", "s", 10, "Asta")
	rootCmd.AddCommand(getCmd)

	updateCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "Tags separados por comas")
	updateCmd.Flags().StringVarP(&title, "title", "T", "", "Título post")
	updateCmd.Flags().StringVarP(&error, "error", "e", "", "Descripción del error")
	updateCmd.Flags().StringVarP(&solution, "solution", "s", "", "Solucion del error")
	updateCmd.Flags().StringVarP(&description, "description", "d", "", "Descripción")
	updateCmd.Flags().StringVarP(&id, "id", "i", "", "Id del registro")
	rootCmd.AddCommand(updateCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
