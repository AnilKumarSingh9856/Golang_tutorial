package main

import (
	"fmt"
	"os"
	"path/filepath"
)


func main(){
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }
	
	// fileInof, err := f.Stat()

	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }
	// fmt.Println("file name ", fileInof.Name())
	// fmt.Println("file or folder ", fileInof.IsDir())
	// fmt.Println("file size ", fileInof.Size())
	// fmt.Println("file permission ", fileInof.Mode())
	// fmt.Println("file modified at ", fileInof.ModTime())


	// // 1. Open File

	// // data, err := os.ReadFile("example.txt")  // not suitable for larger file
	// // if err != nil{
	// // 	panic(err)
	// // }
	// // fmt.Println(string(data))
	

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// // 2. Optimization: Choose a Buffer Size
	// // Your code used 10 bytes. Real apps use 4KB (4096 bytes) or more.
	// // This matches the disk's physical block size for maximum speed.
	// const chunkSize = 1024 // Reading 1KB at a time
	// buf := make([]byte, chunkSize)

	// // 3. The Processing Loop
	// for {
	// 	// Read content into the buffer
	// 	// 'n' = exact number of bytes we successfully read
	// 	n, err := f.Read(buf)

	// 	// 4. Handle EOF (End of File) - The "Exit Condition"
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break // Stop the loop cleanly
	// 		}
	// 		panic(err) // Handle real errors
	// 	}

	// 	// 5. CRITICAL LOGIC: Only read the valid bytes!
	// 	// 'buf' always has size 1024. But if we are at the end of the file,
	// 	// we might only have 5 bytes left.
	// 	// We slice it: buf[:n] ensures we ignore the garbage empty space.
	// 	validData := buf[:n]

	// 	fmt.Printf("Read Block: %q\n", validData)
	// }

	//  reading folder

	// dir, err := os.Open("../")   // "./" -> current dir, "../" -> previous directory
	// if err != nil{
	// 	fmt.Println("There is no folder")
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)   // -1 for all file in current directory

	// for _, file := range fileInfo{
	// 	fmt.Println(file.Name(), file.IsDir())
	// }

	// newfile, err := os.Create("example2.txt")
	// if err != nil{
	// 	fmt.Println("Failed to create new file")
	// 	panic(err)
	// }
	// defer newfile.Close()

	// newfile.WriteString("My name is anil kumar singh")
	// newfile.WriteString(" \nWhat's your name ?")

	// // or 
	// bytes := []byte("\n new way to add value")
	// newfile.Write(bytes)

	// fi, err := os.Open("example2.txt")
	// if err !=  nil {
	// 	fmt.Println("Not find any content inside this file")
	// 	panic(err)
	// }
	// defer fi.Close()

	// buffer := make([]byte, 4000)

	// for {
	// 	n, err := fi.Read(buffer)
		
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		panic(err)
	// 	}

	// 	validCont := buffer[:n]
	// 	fmt.Println(string(validCont))

	// }

	// srcfile, err := os.Open("example2.txt")
	// if err != nil {
	// 	fmt.Println("No file found with this name")
	// 	panic(err)
	// }
	// defer srcfile.Close()

	// dstFile, err := os.Create("example3.txt")
	// if err != nil {
	// 	fmt.Println("Not able to create new file")
	// 	panic(err)
	// }
	// defer dstFile.Close()

	// reader := bufio.NewReader(srcfile)
	// writer := bufio.NewWriter(dstFile)

	// // for {
	// // 	b, err := reader.ReadByte()

	// // 	if err != nil {
	// // 		// 1. Check for EOF first. This is normal behavior, so we just break.
	// // 		if err == io.EOF {
	// // 			break
	// // 		}
	// // 		// If it's a real error (not EOF), then we panic or handle it.
	// // 		fmt.Println("Not able to read the content of source file")
	// // 		panic(err)
	// // 	}

	// // 	e := writer.WriteByte(b)
	// // 	if e != nil {
	// // 		fmt.Println("Not able to write in the destination file")
	// // 		panic(e)
	// // 	}

	// // 	// 2. REMOVED writer.Flush() from here
	// // }

	// // This one line replaces your entire for loop
	// _, err = io.Copy(writer, reader)

	// /* 

	// While the byte-by-byte approach is good for learning, 
	// in a real production app, you would use io.Copy. 
	// It handles the buffering, looping, and EOF logic for you automatically:
	
	// */

	// // 3. Flush once at the very end to write any remaining bytes in memory to disk
	// err = writer.Flush()
	// if err != nil {
	//     panic(err)
	// }

	// fmt.Println("Written to new file successfully")


	// delete a file

	// err := os.Remove("dummy.go")
	// if err != nil{
	// 	panic(err)
	// }

	// fmt.Println("File deleted succesfully")


	// This creates a path compatible with ANY operating system
    // Example: Current Directory -> Parent -> dummy.go
    targetFile := filepath.Join("..", "dummy.go")

    err := os.Remove(targetFile)
    if err != nil {
        // Use fmt.Println to see exactly WHERE it looked
        fmt.Println("Failed to delete at:", targetFile)
        panic(err)
    }

    fmt.Println("File deleted successfully")




}