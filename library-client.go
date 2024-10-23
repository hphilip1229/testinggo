Describe("checking out a book", Ordered, func() {
	var libraryClient *library.Client
	var book *books.Book
	var err error
  
	BeforeAll(func() {
	  libraryClient = library.NewClient()
	  Expect(libraryClient.Connect()).To(Succeed())
	})
  
	It("can fetch a book from a library", func() {
	  book, err = libraryClient.FetchByTitle("Les Miserables")
	  Expect(err).NotTo(HaveOccurred())
	  Expect(book.Title).To(Equal("Les Miserables"))
	})
  
	It("can check out the book", func() {
	  Expect(library.CheckOut(book)).To(Succeed())
	})
  
	It("no longer has the book in stock", func() {
	  book, err = libraryClient.FetchByTitle("Les Miserables")
	  Expect(err).To(MatchError(books.NOT_IN_STOCK))
	  Expect(book).To(BeNil())
	})
  
	AfterAll(func() {
	  Expect(libraryClient.Disconnect()).To(Succeed())
	})
  })