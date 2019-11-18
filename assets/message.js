const legend = (base, {
	color,
	title,
	tickSize = 6,
	width = 320,
	height = 44 + tickSize,
	marginTop = 18,
	marginRight = 0,
	marginBottom = 16 + tickSize,
	marginLeft = 0,
	ticks = width / 64,
	tickFormat,
	tickValues
} = {}) => {
	// https://observablehq.com/@d3/color-legend

	const ramp = (color, n = 256) => {
		const canvas = document.createElement('canvas');
		canvas.width = n;
		canvas.height = 1;

		const context = canvas.getContext("2d");
		for (let i = 0; i < n; ++i) {
			context.fillStyle = color(i / (n - 1));
			context.fillRect(i, 0, 1, 1);
		}

		return canvas;
	};

	const svg = base.append("svg")
		.attr("width", width)
		.attr("height", height)
		.attr("viewBox", [0, 0, width, height])
		.style("overflow", "visible")
		.style("display", "block");

	let x;

	// Continuous
	if (color.interpolator) {
		x = Object.assign(
			color.copy().interpolator(d3.interpolateRound(marginLeft, width - marginRight)),
			{
				range() {
					return [marginLeft, width - marginRight];
				}
			});

		svg.append("image")
			.attr("x", marginLeft)
			.attr("y", marginTop)
			.attr("width", width - marginLeft - marginRight)
			.attr("height", height - marginTop - marginBottom)
			.attr("preserveAspectRatio", "none")
			.attr("xlink:href", ramp(color.interpolator()).toDataURL());

		// scaleSequentialQuantile doesn’t implement ticks or tickFormat.
		if (!x.ticks) {
			if (tickValues === undefined) {
				const n = Math.round(ticks + 1);
				tickValues = d3.range(n).map(i => d3.quantile(color.domain(), i / (n - 1)));
			}
			if (typeof tickFormat !== "function") {
				tickFormat = d3.format(tickFormat === undefined ? ",f" : tickFormat);
			}
		}
	}

	// Discrete
	else if (color.invertExtent) {
		const thresholds
			= color.thresholds ? color.thresholds() // scaleQuantize
				: color.quantiles ? color.quantiles() // scaleQuantile
					: color.domain(); // scaleThreshold

		const thresholdFormat
			= tickFormat === undefined ? d => d
				: typeof tickFormat === "string" ? d3.format(tickFormat)
					: tickFormat;

		x = d3.scaleLinear()
			.domain([-1, color.range().length - 1])
			.rangeRound([marginLeft, width - marginRight]);

		svg.append("g")
			.selectAll("rect")
			.data(color.range())
			.join("rect")
			.attr("x", (d, i) => x(i - 1))
			.attr("y", marginTop)
			.attr("width", (d, i) => x(i) - x(i - 1))
			.attr("height", height - marginTop - marginBottom)
			.attr("fill", d => d);

		tickValues = d3.range(thresholds.length);
		tickFormat = i => thresholdFormat(thresholds[i], i);
	}

	svg.append("g")
		.attr("transform", `translate(0, ${height - marginBottom})`)
		.call(d3.axisBottom(x)
			.ticks(ticks, typeof tickFormat === "string" ? tickFormat : undefined)
			.tickFormat(typeof tickFormat === "function" ? tickFormat : undefined)
			.tickSize(tickSize)
			.tickValues(tickValues))
		//.call(g => g.selectAll(".tick line").attr("y1", marginTop + marginBottom - height))
		.call(g => g.select(".domain").remove())
		.call(g => g.append("text")
			.attr("y", marginTop + marginBottom - height - 6)
			.attr("fill", "currentColor")
			.attr("text-anchor", "start")
			.attr("font-weight", "bold")
			.text(title));
};

!function () {
	// https://observablehq.com/@d3/calendar-view

	const cellSize = 17;
	const width = 954;
	const height = cellSize * 9;
	const countDay = d => d.getUTCDay();

	const parseDate = d3.utcParse("%Y-%m-%d");

	const formatCount = d3.format(",d");
	const formatDate = d3.utcFormat("%-d %b %Y");
	const formatDay = d => "SMTWTFS"[d.getUTCDay()];
	const formatMonth = d3.utcFormat("%b");

	const max = d3.max(heatmap, d => d.MessageCount);
	const color = d3.scaleSequential([0, max], d3.interpolateBuPu).nice();

	legend(d3.select('.heatmap'), {
		color,
		title: "Messages/day",
		width,
		marginLeft: 50,
		ticks: Math.min(max, width / 64),
	});

	const years = d3.nest()
		.key(d => parseDate(d.Date).getUTCFullYear())
		.entries(heatmap)
		.reverse();

	const svg = d3.select('.heatmap')
		.append("svg")
		.attr("viewBox", [0, 0, width, height * years.length])
		.attr("font-family", "sans-serif")
		.attr("font-size", 10);

	const year = svg.selectAll("g")
		.data(years)
		.join("g")
		.attr("transform", (d, i) => `translate(40,${height * i + cellSize * 1.5})`);

	year.append("text")
		.attr("x", -5)
		.attr("y", -5)
		.attr("font-weight", "bold")
		.attr("text-anchor", "end")
		.text(d => d.key);

	year.append("g")
		.attr("text-anchor", "end")
		.selectAll("text")
		.data(d3.range(7).map(i => new Date(1995, 0, i)))
		.join("text")
		.attr("x", -5)
		.attr("y", d => (countDay(d) + 0.5) * cellSize)
		.attr("dy", "0.31em")
		.text(formatDay);

	year.append("g")
		.selectAll("rect")
		.data(d => d.values)
		.join("rect")
		.attr("width", cellSize - 1)
		.attr("height", cellSize - 1)
		.attr("x", d => d3.utcSunday.count(d3.utcYear(parseDate(d.Date)), parseDate(d.Date)) * cellSize + 0.5)
		.attr("y", d => countDay(parseDate(d.Date)) * cellSize + 0.5)
		.attr("fill", d => d.MessageCount === 0 ? '#ebedf0' : color(d.MessageCount))
		.append("title")
		.text(d => `${formatDate(parseDate(d.Date))}: ${formatCount(d.MessageCount)}`);

	year.append("g")
		.selectAll("g")
		.data(d => d3.utcMonths(d3.utcMonth(parseDate(d.values[0].Date)), parseDate(d.values[d.values.length - 1].Date)))
		.join("g")
		.append("text")
		.attr("x", d => d3.utcSunday.count(d3.utcYear(d), d3.utcSunday.ceil(d)) * cellSize + 2)
		.attr("y", -5)
		.text(formatMonth);
}();

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