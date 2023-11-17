# Check Wehrner von Braun's circumlunar probe orbit

![circum lunar vehicle](lunar_probe_bw.png)

[Across the Space Frontier](https://space.nss.org/book-review-across-the-space-frontier/),
Joseph Kaplan, Wernher von Braun, Heinz Haber, Willey Ley, Oscar Schatchter, Fred Whipple,
edited by Cornelius Ryan, Viking Press, 1952.
part of a prose expostion,
and elaboration of von Braun's _The Mars Project_,
the technical appendix of [Project Mars](),

Von Braun describes a lunar excursion,
to be undertaken as a way to ease into interplanetary travel,
and map the landing site for a lunar expedition.

## Vehical Specs

Unfortunately, von Braun did not see fit to numerically describe
the circumlunar vehicle as fully as he did the Three Stage Launch vehicle.
He does write:
"The propulsion unit is the same as that of the Third Stage".

The _PRINCIPAL DATA ON THE THREE_STAGE ROCKET SHIP_ table has these
relevant facts for the Third Stage:

* 220 ton thrust
* 0.77 ton/sec propellant consumption
* 9200 ft/sec exhaust velocity
* 286 sec I<sub>sp</sub>

In consistent SI units:

* Thrust 1.958E6 Kg m/(s<sup>2</sup>
* Propellant consumption 700 kg/s
* Exhaust velocity 2805 m/s

There's a phrase in the book about leaving space station orbit requires "barely 2 minutes" of thrust.

From this, we can back out the propellant consumption of 168,000 kg.
That's one 120 s burn
to leave space station orbit, one to circularize the very eccentric
cislunar orbit.

Von Braun also gives a &#916;V requirement of 2835 m/s to leave the space station's
1075-mile orbit, getting into an orbit that whose apogee is just outside the moon's orbit.

From these facts and the [Rocket Equation](),
we can back out a fully fueled mass of 193652 kg,
and a dry mass of 25652 kg.

## Orbit Diagram

![orbit diagram](circum_lunar_orbit.png)

## A Program for Verifying the Orbit

In accordance with [Gall's Law](http://principles-wiki.net/principles:gall_s_law),
I'm going to start small and work my way up to the hard tasks.
In the process, I intend to learn a little about numerical integration,
and refresh my 40-year-old knowledge of orbital mechanics.

Here's my project plan:

1. Numerically integrate a rocket thrusting for a period.
Double check that the change in velocity matches
   * Wrote a [simulation](csm.go) based on the Apollo Command/Service module,
   since that's very well documented numerically.
2. Numerically integrate a circular orbit. That is,
using (vector) **F** = m **A**, where acceleration **A**
always points to the center of the earth, see if an object with
tangential velocity of 7069.5 meters/sec ends up in a 1075 mile circular orbit.
   * Wrote a [simulation](ss.go) of Wernher von Braun's polar-orbiting
   space station from _Across the Space Frontier]_
3. After one orbit, apply an impulsive velocity change,
determine if **F** = m **A** integration changes to an elliptic orbit.
   * Wrote a simulation of von Braun's lunar excursion vehicle
   doing a 1000 meter/sec impulsive velocity change,
   along with change in mass from expending that much propellant
4. Try numeric integration of a [Hohmann transfer orbit]() works.
This requires two impulsive velocity changes.
5. Try a "barely 2 minute" continuous thrust and concomitant mass change
numerically integrates to a big ellipse that takes the vehicle
out to the radius of the Moon's orbit.
6. Try to put a 3rd body, the Moon, into the simulation of (5).
