// When submit is pressed retrive the table requested
document.getElementById("tableForm").addEventListener("submit", function (e) {
  e.preventDefault();

  fetch("/resourceIndicator/api")
    .then(res => res.json())
     .then(data => {
      const container = document.getElementById("tableDisplay");
      container.innerHTML = ""; // clear any old content

      if (data.error) {
        container.textContent = "Error: " + data.error;
        return;
      }

      // Create the table
      const table = document.createElement("table");
      table.id = "userTable";
      table.className = "table table-dark";

      // Create header
      const header = table.insertRow();
      data.columns.forEach(col => {
        const th = document.createElement("th");
        th.textContent = col;
        header.appendChild(th);
      });

      // Create rows
      data.rows.forEach(row => {
        const tr = table.insertRow();
        row.forEach(cell => {
          const td = tr.insertCell();
          td.textContent = cell;
        });
      });

      container.appendChild(table);
    })
    .catch(err => {
      console.error("Fetch error:", err);
      document.getElementById("tableDisplay").textContent = "An error occurred.";
    });
});

