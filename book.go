var _ = Describe("Books", func() {
	var book *books.Book
  
	BeforeEach(func() {
	  book = &books.Book{
		Title: "Les Miserables",
		Author: "Victor Hugo",
		Pages: 2783,
	  }
	  Expect(book.IsValid()).To(BeTrue())
	})
  
	Describe("Extracting the author's first and last name", func() { ... })
  
	Describe("JSON encoding and decoding", func() {
	  It("survives the round trip", func() {
		encoded, err := book.AsJSON()
		Expect(err).NotTo(HaveOccurred())
  
		decoded, err := books.NewBookFromJSON(encoded)
		Expect(err).NotTo(HaveOccurred())
  
		Expect(decoded).To(Equal(book))
	  })
  
	  Describe("some JSON decoding edge cases", func() {
		var err error
  
		When("the JSON fails to parse", func() {
		  BeforeEach(func() {
			book, err = NewBookFromJSON(`{
			  "title":"Les Miserables",
			  "author":"Victor Hugo",
			  "pages":2783oops
			}`)
		  })
  
		  It("returns a nil book", func() {
			Expect(book).To(BeNil())
		  })
  
		  It("errors", func() {
			Expect(err).To(MatchError(books.ErrInvalidJSON))
		  })
		})
  
		When("the JSON is incomplete", func() {
		  BeforeEach(func() {
			book, err = NewBookFromJSON(`{
			  "title":"Les Miserables",
			  "author":"Victor Hugo",
			}`)
		  })
  
		  It("returns a nil book", func() {
			Expect(book).To(BeNil())
		  })
  
		  It("errors", func() {
			Expect(err).To(MatchError(books.ErrIncompleteJSON))
		  })
		})
	  })
	})
  })