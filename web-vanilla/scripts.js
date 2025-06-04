/**
 * Asks the backend for an image file.
 * @param {string} filename - The name of the image file to request.
 */
async function askForImage(filename) {
    if (!filename) {
        console.error("Filename is required to request an image.");
        return;
    }
    // uploads/1_Interiors/48x48/Room_Builder_48x48.png
    try {
        const response = await fetch(`/uploads/${filename}`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return await response.blob();
    } catch (error) {
        console.error("Error fetching image:", error);
    }
}
