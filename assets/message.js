!function () {
	// https://observablehq.com/@d3/horizontal-bar-chart

	const nameLen = d3.max(senders, d => d.SenderName.length);
	const margin = { top: 30, right: 0, bottom: 10, left: 30 + nameLen * 5 };

	const height = senders.length * 25 + margin.top + margin.bottom;
	const width = d3.select('.senders').node().clientWidth;

	const x = d3.scaleLinear()
		.domain([0, d3.max(senders, d => d.MessageCount)])
		.range([margin.left, width - margin.right]);

	const y = d3.scaleBand()
		.domain(senders.map(d => d.SenderName))
		.range([margin.top, height - margin.bottom])
		.padding(0.1);

	const format = x.tickFormat();

	const svg = d3.select('.senders')
		.append("svg")
		.attr("viewBox", [0, 0, width, height]);

	svg.append("g")
		.selectAll("rect")
		.data(senders)
		.join("rect")
		.attr("fill", (_, i) => d3.schemeSet2[i % d3.schemeSet2.length])
		.attr("x", x(0))
		.attr("y", d => y(d.SenderName))
		.attr("width", d => x(d.MessageCount) - x(0))
		.attr("height", y.bandwidth());

	svg.append("g")
		.attr("fill", "white")
		.attr("text-anchor", "end")
		.style("font", "12px sans-serif")
		.selectAll("text")
		.data(senders)
		.join("text")
		.attr("x", d => x(d.MessageCount) - 4)
		.attr("y", d => y(d.SenderName) + y.bandwidth() / 2)
		.attr("dy", "0.35em")
		.text(d => format(d.MessageCount));

	svg.append("g")
		.attr("transform", `translate(0,${margin.top})`)
		.call(d3.axisTop(x).ticks(width / 80))
		.call(g => g.select(".domain").remove());

	svg.append("g")
		.attr("transform", `translate(${margin.left},0)`)
		.call(d3.axisLeft(y).tickSizeOuter(0));
}();