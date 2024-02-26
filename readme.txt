Explain:

This code processes multipart form data in 32MB chunks (`r.ParseMultipartForm(32 << 20)`) to manage large files efficiently, and it copies files to the destination in chunks (`io.Copy`) instead of loading them entirely into memory. These practices enable the server to handle file uploads up to 10GB while avoiding memory exhaustion.