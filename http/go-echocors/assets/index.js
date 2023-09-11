async function req() {
    let resource = document.getElementById("resource").value
    let options = JSON.parse(document.getElementById("options").value)
    await fetch(resource, options)
}