function tableRowsUpdate(USER_ROW_COUNT) {

    var clusterize = new Clusterize({
        scrollId: 'scrollArea',
        contentId: 'contentArea'
    });

    // part counter
    var part_count = 0;
    // array for users
    var RandomIdArray = [];

    // sending request and adding new rows to table
    function getNewRows(part) {
        // table
        let block = document.getElementById('scrollArea');

        // create request
        var xhr = new XMLHttpRequest();

        // set params
        var params =
            'part=' + encodeURIComponent(part);

        // send request
        xhr.open("POST", '/user/list?' + params, true);
        xhr.send();

        // get responce
        xhr.onload = function() {

            if (xhr.starus == 200) {
                // if bad status
                console.log(xhr.status + ":" + xhr.statusText)
            } else {
                // if success status
                let users = JSON.parse(xhr.response)

                // if new part is not full
                if (users.length == USER_ROW_COUNT) {
                    part_count += 1
                }

                // checking old records
                users = users.filter(user => {
                    return !RandomIdArray.includes(user.RandomId)
                })

                // recording new rows
                users = users.map(user => {
                    RandomIdArray.push(user.RandomId)
                    return `<tr class="table_item">
                        <td>` + user.Username + `</td>
                        <td>` + user.RandomId + `</td>
                        <td>` + user.Register + `</td>
                        </tr>`;
                })
                clusterize.append(users)
            }

        }
    }
    // init rows
    getNewRows(part_count)

    // table scroll div element
    var block = document.getElementById('scrollArea');
    // table 
    var table = document.getElementById("contentArea")
        // checking any scroll event in block
    block.addEventListener("scroll", function() {
        var contentHeight = table.offsetHeight; // heiht of block with borders
        var yOffset = block.scrollTop - 31; // place of scrollbar
        var window_height = block.offsetHeight; // heiht into block 
        var y = yOffset + window_height;

        // if end of table
        if (y >= contentHeight) {
            // send reqest to get new users data
            getNewRows(part_count)
        }
    });
}