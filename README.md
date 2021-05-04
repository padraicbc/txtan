# txtan

[Jaccard similarity](https://en.wikipedia.org/wiki/Jaccard_index) and [Cosine similarity](https://en.wikipedia.org/wiki/Cosine_similarity) implementations in go.


To use just pass in 2 slices of words: 
     
      a := ... slice of words
	  b := .. slice of words

	  an := txtan.Setup(a , b)

	  fmt.Printf("%.2f %.2f\n", an.CosineSimilarity(), an.JaccardSimilarity())
    
No processing like removing stop words, punctuation etc.. That is up to the end user. 
The code tries to use MAXPROCS go routines to split up various parts but it won't be a great deal of use bar for larger bodies of text. 
No benchmarks done either so no doubt plenty improvements for efficency but it works...
