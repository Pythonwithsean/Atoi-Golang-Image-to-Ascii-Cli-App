import { useState } from "react";

function App() {
  async function handleFileUpload(event: React.ChangeEvent<HTMLInputElement>) {
    const fileInput = event.target;
    if (!fileInput.files) {
      return;
    }
    const file = fileInput.files[0];
    const uri = URL.createObjectURL(file);
    setImage(uri);
  }

  async function sendImage() {
    const fileInput = document.querySelector(
      "input[type=file]"
    ) as HTMLInputElement;
    if (!fileInput) {
      return;
    }
    const file = fileInput.files?.[0];
    console.log(file);
    const formData = new FormData();
    formData.append("image", file as Blob);

    const response = await fetch("http://localhost:8080/convert", {
      method: "POST",
      body: formData,
    });
    const data = await response.json();
    const { art } = data;
    setAscii(art);
  }

  const [ascii, setAscii] = useState("/* ASCII art will appear here */");
  const [image, setImage] = useState<string | null>(null);

  return (
    <>
      <header>
        <h1>Atoi</h1>
      </header>
      <main>
        <section id="upload">
          <h2>Image to ASCII Generator</h2>
          <div className="content-container">
            <div className="image-preview">
              {image && (
                <img src={image} alt="Uploaded" className="image-original" />
              )}
              <input type="file" accept="image/*" onChange={handleFileUpload} />
              <button onClick={sendImage}>Generate ASCII</button>
            </div>
            <div className="ascii-preview">
              <pre className="ascii-output">{ascii}</pre>
            </div>
          </div>
        </section>
      </main>
      <footer>
        <p>© 2024 Atoi. All rights reserved.</p>
      </footer>
    </>
  );
}

export default App;
