(() => {
    const dataTable = document.getElementById("data")
    fetch("/employee")
        .then(res => res.json())
        .then(async (array) => {
            for (const index in array) {
                const {firstname,lastname,gender} = array[index]
                const {genderCode} = await fetch("/employee/transform", {
                    method: "POST",
                    headers: {"Content-Type": "application/json"},
                    body: JSON.stringify({firstname,lastname,gender}),
                }).then(res => res.json())
                array[index].genderCode = genderCode
            }

            array.forEach(({firstname,lastname,gender,genderCode}) => {
                dataTable.innerHTML += `<tr><td>${firstname}</td><td>${lastname}</td><td>${gender}</td><td>${genderCode}</td></tr>`
            });
        })
})();