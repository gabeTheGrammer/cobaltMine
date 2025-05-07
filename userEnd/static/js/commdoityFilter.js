document.getElementById("tableForm").addEventListener("submit", function (e) {
  e.preventDefault();

  const tableName = document.getElementById("tableName").value;
  const typeInp = document.getElementById("inputCommodity").value.toLowerCase();

  fetch("/commodityFilter/api", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ 
      commodity: typeInp,
      type: tableName
    })
  })
    .then(res => res.json())
    .then(data => {
      const container = document.getElementById("tableDisplay");
      container.innerHTML = ""; // clear any old content

      if (data.error) {
        container.textContent = "Error: " + data.error;
        return;
      }

      const table = document.createElement("table");
      table.id = "userTable";
      table.className = "table table-dark";

      const header = table.insertRow();
      data.columns.forEach(col => {
        const th = document.createElement("th");
        th.textContent = col;
        header.appendChild(th);
      });

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

