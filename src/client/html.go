package client

const htmlContent = `
<!DOCTYPE html>
<html>
  <head>
    <style>
      body,
      html {
        margin: 0;
        padding: 0;
      }

      .title {
        margin-top: 50px;
        border: 3vh;
      }

      .new {
        margin-top: 80px;
        border: groove;
        padding: 20px;
      }

      #alert-section {
        position: fixed;
        top: 0px;
        right: 0px;
        max-width: 300px;
        z-index: 9999;
        border-radius: 0px;
        margin: 5px;
      }

      .my-custom-scrollbar {
        position: relative;
        height: 400px;
        overflow: auto;
      }

      .table-wrapper-scroll-y {
        margin-top: 30px;
        display: block;
      }

      footer {
        position: fixed;
        left: 0;
        bottom: 0;
        width: 100%;
        background-color: red;
        color: white;
        text-align: center;
      }
      
      #td-value-input, #btn-update, #btn-reject-update {
        display: none;
      }
    </style>
    <link
      href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.1/dist/css/bootstrap.min.css"
      rel="stylesheet"
      integrity="sha384-iYQeCzEYFbKjA/T2uDLTpkwGzCiq6soy8tYaI1GyVh/UjpbCx/TYkiZhlZB6+fzT"
      crossorigin="anonymous"
    />
    <script
      src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.1/dist/js/bootstrap.bundle.min.js"
      integrity="sha384-u1OknCvxWvY5kfmNBILK2hRnQC3Pr17a+RTT6rIHI7NnikvbZlHgTPOOmMi466C8"
      crossorigin="anonymous"
    ></script>
    <script
      src="http://ajax.googleapis.com/ajax/libs/jquery/1.7.1/jquery.min.js"
      type="text/javascript"
    ></script>
  </head>
  <body>
    <div class="container text-center">
      <div class="title">
        <h3>Comandante</h3>
        <p>Manage server configs</p>
      </div>
      <div class="new">
        <form class="row" onsubmit="AddNewConfig(event)">
          <div class="mb-3 col">
            <input
              type="text"
              class="form-control"
              id="key"
              placeholder="Key"
            />
          </div>
          <div class="mb-3 col">
            <input
              type="text"
              class="form-control"
              id="value"
              placeholder="Value"
            />
          </div>
          <div class="row">
            <div class="col">
              <button type="submit" class="btn btn-success">
                Add New Config Variable
              </button>
            </div>
          </div>
        </form>
      </div>
      <div class="table-wrapper-scroll-y my-custom-scrollbar">
        <table class="table table-striped vars">
          <thead class="thead-dark">
            <tr>
              <th scope="col">Key</th>
              <th scope="col">Value</th>
              <th scope="col">Last Updated At</th>
              <th scope="col">Actions</th>
            </tr>
          </thead>
          <tbody>
            {{range .ConfigVariables}}
            <tr>
              <td id="td-key">{{.Key}}</td>
              <td id="td-value">{{.Value}}</td>
              <td id="td-value-input">
                <input
                  type="text"
                  class="form-control"
                  id="td-value-input-input"
                  value="{{.Value}}"
                  placeholder="{{.Value}}"
                /> 
              </td>
              <td>{{.LastUpdatedAt}}</td>
              <td>
                <button type="button" class="btn btn-success" id="btn-pre-update" onclick="PreUpdateConfigVariable(event)">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    fill="currentColor"
                    class="bi bi-pencil"
                    viewBox="0 0 16 16"
                  >
                    <path
                      d="M12.146.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1 0 .708l-10 10a.5.5 0 0 1-.168.11l-5 2a.5.5 0 0 1-.65-.65l2-5a.5.5 0 0 1 .11-.168l10-10zM11.207 2.5 13.5 4.793 14.793 3.5 12.5 1.207 11.207 2.5zm1.586 3L10.5 3.207 4 9.707V10h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.293l6.5-6.5zm-9.761 5.175-.106.106-1.528 3.821 3.821-1.528.106-.106A.5.5 0 0 1 5 12.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.468-.325z"
                    />
                  </svg>
                </button>
                <button type="button" class="btn btn-success" id="btn-update" onclick="UpdateConfigVariable(event)">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-check-lg" viewBox="0 0 16 16">
                    <path d="M12.736 3.97a.733.733 0 0 1 1.047 0c.286.289.29.756.01 1.05L7.88 12.01a.733.733 0 0 1-1.065.02L3.217 8.384a.757.757 0 0 1 0-1.06.733.733 0 0 1 1.047 0l3.052 3.093 5.4-6.425a.247.247 0 0 1 .02-.022Z"/>
                  </svg>
                </button>
                <button type="button" class="btn btn-warning" id="btn-reject-update" onclick="RejectUpdateConfigVariable(event)">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-x-lg" viewBox="0 0 16 16">
                    <path d="M2.146 2.854a.5.5 0 1 1 .708-.708L8 7.293l5.146-5.147a.5.5 0 0 1 .708.708L8.707 8l5.147 5.146a.5.5 0 0 1-.708.708L8 8.707l-5.146 5.147a.5.5 0 0 1-.708-.708L7.293 8 2.146 2.854Z"/>
                  </svg>
                </button>
                <button type="button" class="btn btn-danger" id="btn-remove" onclick="RemoveConfigVariable(event)">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    fill="currentColor"
                    class="bi bi-trash2"
                    viewBox="0 0 16 16"
                  >
                    <path
                      d="M14 3a.702.702 0 0 1-.037.225l-1.684 10.104A2 2 0 0 1 10.305 15H5.694a2 2 0 0 1-1.973-1.671L2.037 3.225A.703.703 0 0 1 2 3c0-1.105 2.686-2 6-2s6 .895 6 2zM3.215 4.207l1.493 8.957a1 1 0 0 0 .986.836h4.612a1 1 0 0 0 .986-.836l1.493-8.957C11.69 4.689 9.954 5 8 5c-1.954 0-3.69-.311-4.785-.793z"
                    />
                  </svg>
                </button>
              </td>
            </tr>
            {{end}}
          </tbody>
        </table>
      </div>
      <footer class="bg-light text-center text-lg-start">
        <div
          class="text-center p-3"
          style="background-color: rgba(0, 0, 0, 0.2)"
        >
          Don't forget to give a star &#x1F62C;
          <div>
            <a
              class="text-dark"
              href="http://github.com/keremdokumaci/comandante"
              >Comandante</a
            >
          </div>
        </div>
      </footer>
    </div>

    <div id="alert-section"></div>
  </body>
  <script>
    function AddNewConfig(e) {
      e.preventDefault();
      $.ajax({
        type: "POST",
        url: window.location.href,
        data: JSON.stringify({
          key: document.getElementById("key").value,
          value: document.getElementById("value").value,
        }),
      })
        .done(function (data) {
          Alert("Config variable was added successfully !", "success");
          setTimeout(() => {
            window.location.reload();
          }, 700);
        })
        .fail(function (err) {
          Alert("Error : " + err.responseText, "danger");
        });
    }
    
    function PreUpdateConfigVariable(e) {
      e.preventDefault();
      $(e.target).closest('tr')[0].children[1].style.display = "none"
      $(e.target).closest('tr')[0].children[2].style.display = "block"
      $(e.target).closest('td').find('#btn-pre-update').removeAttr("style").hide();
      $(e.target).closest('td').find('#btn-update').show();
      $(e.target).closest('td').find('#btn-reject-update').show();
    }
    
    function RejectUpdateConfigVariable(e) {
      e.preventDefault();
      $(e.target).closest('tr')[0].children[1].style.display = ""
      $(e.target).closest('tr')[0].children[2].style.display = "none"
      $(e.target).closest('td').find('#btn-pre-update').show();
      $(e.target).closest('td').find('#btn-update').removeAttr("style").hide();
      $(e.target).closest('td').find('#btn-reject-update').removeAttr("style").hide();
      
    }
    
    function UpdateConfigVariable(e) {
      e.preventDefault();
      const key = $($(e.target).closest('tr')).find('#td-key')[0].textContent;
      const newValue = $($(e.target).closest('tr')).find('#td-value-input-input')[0].value;
      if(!newValue) {
        Alert("Value must be filled", "danger")
        return
      }
      $.ajax({
        type: "PUT",
        url: window.location.href,
        data: JSON.stringify({
          key: key,
          value: newValue
        }),
      })
        .done(function (data) {
          Alert("Config variable was updated successfully !", "success");
          setTimeout(() => {
            window.location.reload();
          }, 700);
        })
        .fail(function (err) {
          Alert("Error : " + err.responseText, "danger");
        });
    }

    function RemoveConfigVariable(e) {
      e.preventDefault();
      $.ajax({
        type: "DELETE",
        url: window.location.href,
        data: JSON.stringify({
          key: $(e.target).closest('td')[0].parentElement.children[0].textContent,
        }),
      })
        .done(function (data) {
          Alert("Config variable was removed successfully !", "success");
          setTimeout(() => {
            window.location.reload();
          }, 700);
        })
        .fail(function (err) {
          Alert("Error : " + err.responseText, "danger");
        });
    }

    function Alert(message, type) {
      const alertSection = document.getElementById("alert-section");
      const wrapper = document.createElement("div");
      
      const alertDiv = '<div class="alert alert-'+type+' alert-dismissible role="alert">'
      const messageDiv = '<div>'+message+'</div>'
      const closeButton = '<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>'
      const closeDiv = '</div>'
      
      wrapper.innerHTML = [
        alertDiv,
        messageDiv,
        closeButton,
        closeDiv,
      ].join("");
      
      alertSection.append(wrapper);
    }

    function OpenRepository() {
      window.location.replace("http://github.com/keremdokumaci/comandante");
    }
  </script>
</html>
`
