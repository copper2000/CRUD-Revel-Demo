
function Create() {

    var name = $('#name').val();
    var type = $('#type').val();
    var engine = $('#engine').val();

    if (name.length > 0) {
        $.ajax({
            type: "POST",
            url: "/cars/create",
            contentType: "application/json",
            data: JSON.stringify({name: name, type: type, engine: engine}),
            dataType: "json",
            success: function (data) {
                window.location.replace("/cars")
            },
            error: function (data) {

            }
        })
    }
}

function Update() {
    var id = $('#id').val()
    var ids = parseInt(id)
    var name = $('#name').val();
    var type = $('#type').val();
    var engine = $('#engine').val();
    if (ids > 0) {
        $.ajax({
            type: "POST",
            url: "/cars/update",
            contentType: "application/json",
            data: JSON.stringify({id: ids, name: name, type: type, engine: engine}),
            dataType: "json",
            success: function (data) {
                window.location.replace("/cars")
            },
            error: function (data) {

            }
        })
    }
}

function Delete(id) {
    var c = confirm("Do you want to delete this car?")
    if (c == true) {
        if (id > 0) {
            $.ajax({
                type: "POST",
                url: "/cars/delete",
                contentType: "application/json",
                data: JSON.stringify({id: id}),
                dataType: "json",
                success: function (data) {
                    window.location.replace("/cars")
                },
                error: function (data) {

                }
            })
        }
    }

}