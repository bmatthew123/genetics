package genetics

type Genepool interface {

	// Create the initial population.
	Populate()

	// Add a gene to the population
	AddGenes(...Gene)

	// Return the best solution so far.
	BSSF() Gene

	// Prunes out the weakest genes, but could do nothing if necessary.
	Prune()

	// Select two genes to combine.
	SelectGenes() (Gene, Gene)

	// Function to determine if stopping criteria is met
	StoppingCriteriaMet() bool
}

type Gene interface {

	// Combines the gene with the given Gene to produce more Gene(s).
	Crossover(Gene) []Gene

	// Calculate the fitness score of the gene.
	Fitness() int

	// Perform mutations on the Gene.
	Mutate()
}

func Evolve(gp Genepool) Gene {
	gp.Populate()
	for !gp.StoppingCriteriaMet() {
		g1, g2 := gp.SelectGenes()
		genes := g1.Crossover(g2)
		for _, g := range genes {
			g.Mutate()
		}
		gp.AddGenes(genes...)
		gp.Prune()
	}
	return gp.BSSF()
}
