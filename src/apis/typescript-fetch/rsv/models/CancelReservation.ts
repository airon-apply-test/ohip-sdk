/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CancelReservationType } from './CancelReservationType';
import {
    CancelReservationTypeFromJSON,
    CancelReservationTypeFromJSONTyped,
    CancelReservationTypeToJSON,
} from './CancelReservationType';
import type { CancellationReasonType } from './CancellationReasonType';
import {
    CancellationReasonTypeFromJSON,
    CancellationReasonTypeFromJSONTyped,
    CancellationReasonTypeToJSON,
} from './CancellationReasonType';
import type { ChannelResvRQInfoType } from './ChannelResvRQInfoType';
import {
    ChannelResvRQInfoTypeFromJSON,
    ChannelResvRQInfoTypeFromJSONTyped,
    ChannelResvRQInfoTypeToJSON,
} from './ChannelResvRQInfoType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { RateChangeInstructionType } from './RateChangeInstructionType';
import {
    RateChangeInstructionTypeFromJSON,
    RateChangeInstructionTypeFromJSONTyped,
    RateChangeInstructionTypeToJSON,
} from './RateChangeInstructionType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Request object for canceling reservations.
 * @export
 * @interface CancelReservation
 */
export interface CancelReservation {
    /**
     * 
     * @type {RateChangeInstructionType}
     * @memberof CancelReservation
     */
    rateChangeInstruction?: RateChangeInstructionType;
    /**
     * 
     * @type {CancellationReasonType}
     * @memberof CancelReservation
     */
    reason?: CancellationReasonType;
    /**
     * Information on the reservation that is to be canceled.
     * @type {Array<CancelReservationType>}
     * @memberof CancelReservation
     */
    reservations?: Array<CancelReservationType>;
    /**
     * 
     * @type {ChannelResvRQInfoType}
     * @memberof CancelReservation
     */
    channelInformation?: ChannelResvRQInfoType;
    /**
     * Indicator if the request is a verification on whether the reservation can be canceled.
     * @type {boolean}
     * @memberof CancelReservation
     */
    verificationOnly?: boolean;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof CancelReservation
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CancelReservation
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the CancelReservation interface.
 */
export function instanceOfCancelReservation(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CancelReservationFromJSON(json: any): CancelReservation {
    return CancelReservationFromJSONTyped(json, false);
}

export function CancelReservationFromJSONTyped(json: any, ignoreDiscriminator: boolean): CancelReservation {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'rateChangeInstruction': !exists(json, 'rateChangeInstruction') ? undefined : RateChangeInstructionTypeFromJSON(json['rateChangeInstruction']),
        'reason': !exists(json, 'reason') ? undefined : CancellationReasonTypeFromJSON(json['reason']),
        'reservations': !exists(json, 'reservations') ? undefined : ((json['reservations'] as Array<any>).map(CancelReservationTypeFromJSON)),
        'channelInformation': !exists(json, 'channelInformation') ? undefined : ChannelResvRQInfoTypeFromJSON(json['channelInformation']),
        'verificationOnly': !exists(json, 'verificationOnly') ? undefined : json['verificationOnly'],
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function CancelReservationToJSON(value?: CancelReservation | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'rateChangeInstruction': RateChangeInstructionTypeToJSON(value.rateChangeInstruction),
        'reason': CancellationReasonTypeToJSON(value.reason),
        'reservations': value.reservations === undefined ? undefined : ((value.reservations as Array<any>).map(CancelReservationTypeToJSON)),
        'channelInformation': ChannelResvRQInfoTypeToJSON(value.channelInformation),
        'verificationOnly': value.verificationOnly,
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

