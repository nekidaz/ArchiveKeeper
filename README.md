# ArchiveKeeper

## Description

The File Archiver project is a simple web application written in Go that allows users to upload files, archive them into a ZIP file, and download the generated archive. It utilizes the Gin web framework and the archiver library for file handling and archiving operations.

## Features

- Upload files: Users can select and upload files through the web interface.
- Archive files: The uploaded files are archived into a ZIP file using the archiver library.
- Download archives: Users can download the generated ZIP archives containing their uploaded files.

## Installation

1. Install Go: Ensure that Go is installed on your machine. You can download and install Go from the official website: https://golang.org/

2. Clone the repository: Use `git clone` to clone the repository to your local machine.

3. Set up Go module: Open a terminal or command prompt, navigate to the project directory, and run `go mod init <module-name>` to initialize the Go module for the project.

4. Install dependencies: Run `go mod install` to download and install the required dependencies specified in the go.mod file.

## Usage

1. Start the application: In the terminal or command prompt, navigate to the project directory and run `go run main.go` to start the web application.

2. Access the application: Open a web browser and visit `http://localhost:8080` to access the main page of the application.

3. Upload files: On the main page, click the "Choose File" button to select and upload files. Files should not exceed 100 MB in size.

4. Archive files: After uploading the files, they will be automatically archived into a ZIP file.

5. Download archives: Once the archiving process is complete, a download link will appear. Click the link to download the generated ZIP archive containing the uploaded files.


