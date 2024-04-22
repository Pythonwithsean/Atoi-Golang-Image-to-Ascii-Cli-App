Atoi-Golang-Image-to-Ascii-Cli-App

Algorithm 

Calculate Aspect Ratio: Determine the aspect ratio of the original image and scale it down to fit within a reasonable width for ASCII art. You can calculate the aspect ratio as aspectRatio = width / height.
Resize Image: Resize the original image to match the desired width while maintaining the aspect ratio. You can use image processing libraries like Go's image package to resize the image.
Map Brightness to ASCII Characters: As you've done, calculate the brightness for each pixel using the luminance formula and map it to appropriate ASCII characters.
Print ASCII Art: Print the ASCII characters row by row, ensuring that each ASCII character represents an appropriate region of the resized image. You may need to adjust the brightness thresholds and ASCII characters to achieve the desired visual effect.
Adjust Character Density: Since ASCII characters are taller than they are wide, you may need to use fewer characters vertically compared to horizontally to maintain the aspect ratio. You can achieve this by skipping rows or columns as needed while iterating over the image pixels.
