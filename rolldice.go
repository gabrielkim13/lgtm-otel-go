package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	tracer  = otel.Tracer("rolldice")
	meter   = otel.Meter("rolldice")
	rollCnt metric.Int64Counter
)

func init() {
	var err error
	rollCnt, err = meter.Int64Counter("dice.rolls",
		metric.WithDescription("The number of rolls by roll value"),
		metric.WithUnit("{roll}"))
	if err != nil {
		panic(err)
	}
}

func rolldice(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "roll")
	defer span.End()

	sleep(ctx)

	roll := 1 + rand.Intn(6)

	span.AddEvent("Rolled dice")

	log.Printf("Rolled a %d\n", roll)

	if roll == 6 {
		log.Println("Lucky!")
	}

	rollValueAttr := attribute.Int("roll.value", roll)
	span.SetAttributes(rollValueAttr)
	rollCnt.Add(ctx, 1, metric.WithAttributes(rollValueAttr))

	resp := strconv.Itoa(roll) + "\n"
	if _, err := io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}

func sleep(ctx context.Context) {
	_, span := tracer.Start(ctx, "sleep")
	defer span.End()

	span.AddEvent("Sleeping")

	sleep100ms(ctx)
	sleep200ms(ctx)
	sleep300ms(ctx)
}

func sleep100ms(ctx context.Context) {
	_, span := tracer.Start(ctx, "sleep100ms")
	defer span.End()

	span.AddEvent("Sleeping for 100ms")

	time.Sleep(100 * time.Millisecond)
}

func sleep200ms(ctx context.Context) {
	_, span := tracer.Start(ctx, "sleep200ms")
	defer span.End()

	span.AddEvent("Sleeping for 200ms")

	time.Sleep(200 * time.Millisecond)
}

func sleep300ms(ctx context.Context) {
	_, span := tracer.Start(ctx, "sleep300ms")
	defer span.End()

	span.AddEvent("Sleeping for 300ms")

	time.Sleep(300 * time.Millisecond)
}
