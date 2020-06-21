package metric

import "github.com/prometheus/client_golang/prometheus"

/*
input.sensitivity: high
input.transfer.high: 253
input.transfer.low: 200
input.voltage: 228.9
output.current: 0.00
output.frequency: 50.0
output.voltage: 228.9
output.voltage.nominal: 230.0



DATE     : 2020-06-18 16:56:59 +0300
HOSTNAME : home-nas
VERSION  : 3.14.14 (31 May 2016) debian
UPSNAME  : HomeSrv
CABLE    : USB Cable
DRIVER   : USB UPS Driver
UPSMODE  : Stand Alone
STARTTIME: 2020-06-18 02:00:02 +0300
MODEL    : Back-UPS CS 500
STATUS   : ONLINE

LINEV    : 226.0 Volts
LOADPCT  : 7.0 Percent
BCHARGE  : 100.0 Percent
TIMELEFT : 64.5 Minutes
MBATTCHG : 1 Percent
MINTIMEL : -1 Minutes
MAXTIME  : 0 Seconds
OUTPUTV  : 230.0 Volts
LINEFREQ : 50.0 Hz
NOMPOWER : 300 Watts
NOMINV   : 230 Volts
NOMOUTV  : 230 Volts
TONBATT  : 0 Seconds
CUMONBATT: 0 Seconds
XOFFBATT : N/A  date time
XONBATT     date time
SENSE    : Low|High
LOTRANS  : 180.0 Volts
HITRANS  : 260.0 Volts
BATTV    : 13.7 Volts
NOMBATTV : 12.0 Volts
ITEMP    : 29.2 C
BATTDATE : 2018-11-29
MANDATE  : 2018-01-09
DSHUTD   : 180 Seconds
DWAKE    : 0 Seconds
**DLOWBATT**
    The remaining runtime below which the UPS
    sends the low battery signal. At this point apcupsd will force an
	immediate emergency shutdown.
NUMXFERS : 0
RETPCT   : 0.0 Percent
**HUMIDITY**
	The humidity as measured by the UPS.

**AMBTEMP**
    The ambient temperature as measured by the UPS.

**EXTBATTS**
    The number of external batteries as
    defined by the user. A correct number here helps the UPS compute
    the remaining runtime more accurately.

**BADBATTS**
	The number of bad battery packs.

ALARMDEL : No alarm|30 Seconds
SELFTEST : NO
STATFLAG : 0x05000008
STESTI   : None|14 days



LASTXFER : Unacceptable line voltage changes
| No transfers since turnon
| Automatic or explicit self test



SERIALNO : 4B1802P05216
FIRMWARE : 808.q10 .I USB FW:q

*/

// Metrics declare
var Metrics = []*Metric{

	// Input
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_sensitivity",
			Help: `**SENSE** The sensitivity level of the UPS to line voltage fluctuations.
					Unknown=0, Low=1, Medium=2, High=3, 'Auto Adjust'=4`,
		}),
		Handler: DefaultHandler{
			ApcKey: "SENSE",
			Mapper: DictMapper{
				Dict: map[string]float64{
					"Unknown":     0,
					"Low":         1,
					"Medium":      2,
					"High":        3,
					"Auto Adjust": 4,
				},
			},
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_frequency",
			Help: "**LINEFREQ** Line frequency in hertz as given by the UPS.",
		}),
		Handler: NewDefaultHandler("LINEFREQ"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage",
			Help: "**LINEV** The current line voltage as returned by the UPS.",
		}),
		Handler: NewDefaultHandler("LINEV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_min",
			Help: "**MINLINEV** The minimum line voltage since the UPS was started, as returned by the UPS",
		}),
		Handler: NewDefaultHandler("MINLINEV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_max",
			Help: "**MAXLINEV** The maximum line voltage since the UPS was started, as reported by the UPS",
		}),
		Handler: NewDefaultHandler("MAXLINEV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_nominal",
			Help: "**NOMINV** The input voltage that the UPS is configured to expect.",
		}),
		Handler: NewDefaultHandler("NOMINV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_transfer_low",
			Help: "**LOTRANS** The line voltage below which the UPS will switch to batteries.",
		}),
		Handler: NewDefaultHandler("LOTRANS"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_input_voltage_transfer_high",
			Help: "**HITRANS** The line voltage above which the UPS will switch to batteries.",
		}),
		Handler: NewDefaultHandler("HITRANS"),
	},

	// Output
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_load",
			Help: "**LOADPCT** The percentage of load capacity as estimated by the UPS.",
		}),
		Handler: NewDefaultHandler("LOADPCT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_power_nominal",
			Help: "**NOMPOWER** The maximum power in Watts that the UPS is designed to supply.",
		}),
		Handler: NewDefaultHandler("NOMPOWER"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_voltage",
			Help: "**OUTPUTV** The voltage the UPS is supplying to your equipment",
		}),
		Handler: NewDefaultHandler("OUTPUTV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_output_voltage_nominal",
			Help: "**NOMOUTV** The output voltage that the UPS will attempt to supply when on battery power.",
		}),
		Handler: NewDefaultHandler("NOMOUTV"),
	},

	// Battery
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_charge",
			Help: "**BCHARGE** The percentage charge on the batteries.",
		}),
		Handler: NewDefaultHandler("BCHARGE"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_voltage",
			Help: "**BATTV** Battery voltage as supplied by the UPS.",
		}),
		Handler: NewDefaultHandler("BATTV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_voltage_nominal",
			Help: "**NOMBATTV** The nominal battery voltage.",
		}),
		Handler: NewDefaultHandler("NOMBATTV"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_external",
			Help: "**EXTBATTS** The number of external batteries as defined by the user. A correct number here helps the UPS compute the remaining runtime more accurately.",
		}),
		Handler: NewDefaultHandler("EXTBATTS"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_bad",
			Help: "**BADBATTS** The number of bad battery packs.",
		}),
		Handler: NewDefaultHandler("BADBATTS"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_battery_replaced_timestamp",
			Help: "**BATTDATE** The date that batteries were last replaced.",
		}),
		Handler: NewDefaultHandler("BATTDATE"),
	},

	// Ups
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_manafactured_timestamp",
			Help: "**MANDATE** The date the UPS was manufactured.",
		}),
		Handler: NewDefaultHandler("MANDATE"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_status_flag",
			Help: "**STATFLAG** Status flag. English version is given by STATUS.",
		}),
		Handler: NewDefaultHandler("STATFLAG"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_dip_switch_flag",
			Help: "**DIPSW** The current dip switch settings on UPSes that have them.",
		}),
		Handler: NewDefaultHandler("DIPSW"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_reg1",
			Help: "**REG1** The value from the UPS fault register 1.",
		}),
		Handler: NewDefaultHandler("REG1"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_reg2",
			Help: "**REG2** The value from the UPS fault register 2.",
		}),
		Handler: NewDefaultHandler("REG2"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_reg3",
			Help: "**REG3** The value from the UPS fault register 3.",
		}),
		Handler: NewDefaultHandler("REG3"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_timeleft",
			Help: "**TIMELEFT** (seconds) The remaining runtime left on batteries as estimated by the UPS.",
		}),
		Handler: NewDefaultHandler("TIMELEFT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_timeleft_low_battery",
			Help: "**DLOWBATT** (seconds) The remaining runtime below which the UPS sends the low battery signal. At this point apcupsd will force an immediate emergency shutdown.",
		}),
		Handler: NewDefaultHandler("DLOWBATT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery",
			Help: "**NUMXFERS** The number of transfers to batteries since apcupsd startup.",
		}),
		Handler: NewDefaultHandler("NUMXFERS"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_time",
			Help: "**TONBATT** Time in seconds currently on batteries",
		}),
		Handler: NewDefaultHandler("TONBATT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_time_cumulative",
			Help: "Total (cumulative) time on batteries in seconds since apcupsd startup.",
		}),
		Handler: NewDefaultHandler("CUMONBATT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_onbattery_timestamp",
			Help: "**XONBATT** Time and date of last transfer to batteries",
		}),
		Handler: NewDefaultHandler("XONBATT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_transfer_offbattery_timestamp",
			Help: "Time and date of last transfer from batteries",
		}),
		Handler: NewDefaultHandler("XOFFBATT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_turnon_delay",
			Help: "**DWAKE** (seconds) The amount of time the UPS will wait before restoring power to your equipment after a power off condition when the power is restored.",
		}),
		Handler: NewDefaultHandler("DWAKE"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_turnon_battery_min",
			Help: "	**RETPCT** The percentage charge that the batteries must have after a power off condition before the UPS will restore power to your equipment.",
		}),
		Handler: NewDefaultHandler("RETPCT"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_turnoff_delay",
			Help: "**DSHUTD** (seconds) The grace delay that the UPS gives after receiving a power down command from apcupsd before it powers off your equipment.",
		}),
		Handler: NewDefaultHandler("DSHUTD"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_temp_internal",
			Help: "**ITEMP** (Celsius) Internal UPS temperature as supplied by the UPS.",
		}),
		Handler: NewDefaultHandler("ITEMP"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_temp_ambient",
			Help: "**AMBTEMP** The ambient temperature as measured by the UPS.",
		}),
		Handler: NewDefaultHandler("AMBTEMP"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_humidity",
			Help: "**HUMIDITY** The humidity as measured by the UPS.",
		}),
		Handler: NewDefaultHandler("HUMIDITY"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_alarm_mode",
			Help: "**ALARMDEL** The delay period for the UPS alarm.\n" +
				"'No alarm'=0, 'Always'=1, '5 Seconds'=2, '30 Seconds'=3, 'Low Battery'=4",
		}),
		Handler: DefaultHandler{
			ApcKey: "ALARMDEL",
			Mapper: DictMapper{
				Dict: map[string]float64{
					"No alarm":    0,
					"Always":      1,
					"5 Seconds":   2,
					"5":           2,
					"30 Seconds":  3,
					"30":          3,
					"Low Battery": 4,
				},
			},
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_selftest_result",
			Help: "**SELFTEST** The results of the last self test, and may have the following values.\n" +
				"NO=0 No results i.e. no self test performed in the last 5 minutes,\n" +
				"OK=1 self test indicates good battery,\n" +
				"BT=2 self test failed due to insufficient battery capacity,\n" +
				"NG=3 self test failed due to overload",
		}),
		Handler: DefaultHandler{
			ApcKey: "SELFTEST",
			Mapper: DictMapper{
				Dict: map[string]float64{
					"NO": 0,
					"OK": 1,
					"BT": 2,
					"NG": 3,
				},
			},
		},
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_ups_selftest_interval",
			Help: "**STESTI** The interval in seconds between automatic self tests.",
		}),
		Handler: NewDefaultHandler("STESTI"),
	},

	// Shutdown
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_battery_min",
			Help: "**MBATTCHG** If the battery charge percentage (BCHARGE) drops below this value, apcupsd will  shutdown your system.",
		}),
		Handler: NewDefaultHandler("MBATTCHG"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_timeleft_min",
			Help: "**MINTIMEL** (seconds) apcupsd will shutdown your system if the remaining runtime equals or is below this point.",
		}),
		Handler: NewDefaultHandler("MINTIMEL"),
	},
	{
		Collector: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "apcupsd_shutdown_onbattery_time_max",
			Help: "**MAXTIME** (seconds) apcupsd will shutdown your system if the time on batteries exceeds this value. A value of zero disables the feature.",
		}),
		Handler: NewDefaultHandler("MAXTIME"),
	},
}

// RegisterPermanents registering permanents
func RegisterPermanents() {
	for _, m := range Metrics {
		if m.IsPermanent {
			m.Register()
		}
	}
}