function create() {
    var golangVar = "{{.Image}}";
    console.log(golangVar);
    const normal = /https:\/\/groupietrackers.herokuapp.com\/api\/images\/\w*.jpeg/g;
    var img = golangVar.match(normal);
    console.log(img);

    for (i = 0; i <= 55; i++) {
        var div1 = document.createElement('div');
        div1.setAttribute('class', 'card airforce light-1 shadow-1 dropshadow-5')

        var div2 = document.createElement('div');
        div2.setAttribute('class', 'card-image')


        var divimage = document.createElement('img');
        divimage.setAttribute('class', 'img-responsive')
        divimage.setAttribute('id', 'img')
        divimage.setAttribute('src', img[i])
        divimage.setAttribute('alt', 'Music')

        var div3 = document.createElement('div');
        var txt3 = document.createTextNode("Card Header")
        div3.setAttribute('class', 'card-header')
        div3.appendChild(txt3)

        var div4 = document.createElement('div');
        var txt4 = document.createTextNode("Lorem ipsum dolor, sit amet consectetur adipisicing elit. Eligendi suscipit harum repellat architecto unde vel numquam rem doloribus maiores deserunt tenetur labore, aut, adipisci sit, sequi nihil voluptas commodi? Velit?")
        div4.setAttribute('class', 'card-content')
        div4.appendChild(txt4)

        var div5 = document.createElement('div');
        var txt5 = document.createTextNode("Card Footer")
        div5.setAttribute('class', 'card-footer')
        div5.appendChild(txt5)
        var parent = document.getElementById("start")

        document.body.appendChild(div1);
        parent.appendChild(div1)
        div1.appendChild(div2)
        div1.appendChild(div3)
        div1.appendChild(div4)
        div1.appendChild(div5)
        div2.appendChild(divimage)
    }
}

function deletes() {
    const myNode = document.getElementById("start")
    myNode.innerHTML = ""
}

function displays(a) {
    var x = document.getElementById(a);
    x.style.display = "block"
}

function displaysnot(a) {
    var x = document.getElementById(a);
    x.style.display = "none"
}