// Reason for the existence of this package is to be able to perform monero operations independent
// of monero daemon or wallet RPC connections.
// Second is to simplify the protocol implementation as well as the usage of third party
// libraries that should assist us in the process.
// Third to have some documentation when calling and using functions dealing with monero.
//
// After searching for golang monero related libraries there is lack of choice(or nonexistent).
// with the exception of https://github.com/paxosglobal/moneroutil  the only one(original) implementation
// from which this library was partly inspired from.
// https://github.com/exantech/moneroutil is probably the second one with the most improvements
// forked from moneroutil(paxosglobal) but seems abandoned as well.
// Other golang tool that I have discovered are also based on moneroutil(paxosglobal) but
// the implementations were a mess in library usage and I think personally less is more.
// and that is why this library here.
// It implements the bare minimum to be useful and compatible with the monero protocol
// but with freedom in the way you obtain and process data.
//
// resources that helped us so far
//
// https://www.getmonero.org/library/Zero-to-Monero-2-0-0.pdf
// https://masteringmonero.com/free-download.html
// https://www.getmonero.org/library/MoneroAddressesCheatsheet20201206.pdf
// and a lot of monero.stackexchange.com questions
//
// TODO implement RingCT related functions
//
// if you would like to test it and show your appreciation try sending here
// 88nYxA5xZEfLDuTPiBXZuzMRKFzHsR6JJSnBoNkJb9rF16KZxtYzFHJcZoaFKAbeUxXtPUQgjZ6zj7y5WBiP5c8vCXP5r8N
package gonero
