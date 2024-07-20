describe('Post test', () => {
  it('Fills the form', () => {
    cy.visit('/')
    
    // name
    cy.get('#name').type('fake')
    cy.get('#name').should('have.value', 'fake')
    
    //age
    cy.get('#age').type('19')
    cy.get('#age').should('have.value', '19')
    
    // email
    cy.get('#email').type('fake@email.com')
    cy.get('#email').should('have.value', 'fake@email.com')
    
    // description
    cy.get('#description').type('I like donuts')
    cy.get('#description').should('have.value', 'I like donuts')
    cy.get("form").submit()
    cy.get("#post-button").click()
  })


})
