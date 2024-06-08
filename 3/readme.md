Using GIN Framework demonstrate custom binding request(At least give 5 field example with
including boolean, pointers, int, string, runes.

To test this example, you can send a POST request with form data containing the required fields (string_field, int_field, bool_field, rune_field, pointer_field) to the /custom-request endpoint. The server will bind the request data to the CustomRequest struct and return the JSON representation of the bound object