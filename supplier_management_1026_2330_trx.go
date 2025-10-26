// 代码生成时间: 2025-10-26 23:30:36
package controllers

import (
    "encoding/json"
    "github.com/revel/revel"
    "app/models" // Assuming models package contains Supplier model
)

type Supplier struct {
    Id   int    "json:"id""
    Name string "json:"name""
    // Other supplier fields...
}

type Suppliers struct {
    *revel.Controller
}

// Index action renders the supplier list page
func (c Suppliers) Index() revel.Result {
    // Retrieve all suppliers from the database
    suppliers, err := models.GetAllSuppliers()
    if err != nil {
        return c.RenderError(err)
    }
    // Render the suppliers template with the supplier list
    return c.Render(suppliers)
}

// Add action handles the form submission for adding a new supplier
func (c Suppliers) Add() revel.Result {
    var supplier models.Supplier
    // Decode the JSON request body into the supplier struct
    err := json.Unmarshal(c.Params.Form.Body, &supplier)
    if err != nil {
        return c.RenderError(err)
    }
    // Validate supplier details
    if supplier.Name == "" {
        return c.RenderError(errors.New("Supplier name is required"))
    }
    // Save the supplier to the database
    err = models.AddSupplier(&supplier)
    if err != nil {
        return c.RenderError(err)
    }
    // Redirect to the supplier list page
    return c.Redirect(Suppliers.Index)
}

// Edit action handles the editing of an existing supplier
func (c Suppliers) Edit(id int) revel.Result {
    // Retrieve the supplier by ID from the database
    supplier, err := models.GetSupplierById(id)
    if err != nil {
        return c.RenderError(err)
    }
    // Render the edit supplier template with the supplier details
    return c.Render(supplier)
}

// Update action handles the form submission for updating an existing supplier
func (c Suppliers) Update() revel.Result {
    var supplier models.Supplier
    // Decode the JSON request body into the supplier struct
    err := json.Unmarshal(c.Params.Form.Body, &supplier)
    if err != nil {
        return c.RenderError(err)
    }
    // Validate supplier details
    if supplier.Name == "" {
        return c.RenderError(errors.New("Supplier name is required"))
    }
    // Update the supplier in the database
    err = models.UpdateSupplier(&supplier)
    if err != nil {
        return c.RenderError(err)
    }
    // Redirect to the supplier list page
    return c.Redirect(Suppliers.Index)
}

// Delete action handles the deletion of a supplier
func (c Suppliers) Delete(id int) revel.Result {
    // Delete the supplier from the database by ID
    err := models.DeleteSupplier(id)
    if err != nil {
        return c.RenderError(err)
    }
    // Redirect to the supplier list page
    return c.Redirect(Suppliers.Index)
}
