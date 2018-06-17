{{ template "main.tpl" }}

{{ define "pagetitle" }}
  Recieved SMS
{{ end }}

{{ define "breadcrumbs" }}
  Recieved SMS
{{ end }}


{{ define "pagecontent"}}
<table id="smslogs" class="table table-hover table-responsive-lg table-striped">
<thead>
    <tr>
        <th>Sr</th>
        <th>Phone Number</th>
        <th>Message</th>
        <th>Date Time</th>
        <th>Action</th>
    </tr>
</thead>
<tbody>
  <tr>
    <td>1</td>
    <td>1234567890</td>
    <td>Test</td>
    <td>2018-05-04 5:00:12 PM </td>
    <td>Delete</td>
  </tr>
  <tr>
    <td>1</td>
    <td>1234567890</td>
    <td>Test</td>
    <td>2018-05-04 5:00:12 PM </td>
    <td>Delete</td>
  </tr>
  <tr>
    <td>1</td>
    <td>1234567890</td>
    <td>Test</td>
    <td>2018-05-04 5:00:12 PM </td>
    <td>Delete</td>
  </tr>
  <tr>
    <td>1</td>
    <td>1234567890</td>
    <td>Test</td>
    <td>2018-05-04 5:00:12 PM </td>
    <td>Delete</td>
  </tr>
  <tr>
    <td>1</td>
    <td>1234567890</td>
    <td>Test</td>
    <td>2018-05-04 5:00:12 PM </td>
    <td>Delete</td>
  </tr>
</tbody>
</table>
{{ end }}

{{ define "js" }}
<script>
  var sendsms = function(){
    alert("SEND SMS CALLED");
  }

  $(document).ready(function() {
    $('#smslogs').DataTable();
  });
</script>
{{ end }}
