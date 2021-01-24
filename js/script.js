
function showmessage(event) {
    event.preventDefault()
    document.querySelector("#returnDiv").innerHTML = ("Output " + document.querySelector("#name").value + " " + document.querySelector("#message").value)

}
