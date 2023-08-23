/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * Information regarding the message sent to guest.
 * @export
 * @interface QueueTextInfoType
 */
export interface QueueTextInfoType {
    /**
     * Time the text was last sent to the user.
     * @type {Date}
     * @memberof QueueTextInfoType
     */
    sentTime?: Date;
    /**
     * User name of the user who sent message.
     * @type {string}
     * @memberof QueueTextInfoType
     */
    sentBy?: string;
}

/**
 * Check if a given object implements the QueueTextInfoType interface.
 */
export function instanceOfQueueTextInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function QueueTextInfoTypeFromJSON(json: any): QueueTextInfoType {
    return QueueTextInfoTypeFromJSONTyped(json, false);
}

export function QueueTextInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): QueueTextInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'sentTime': !exists(json, 'sentTime') ? undefined : (new Date(json['sentTime'])),
        'sentBy': !exists(json, 'sentBy') ? undefined : json['sentBy'],
    };
}

export function QueueTextInfoTypeToJSON(value?: QueueTextInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'sentTime': value.sentTime === undefined ? undefined : (value.sentTime.toISOString()),
        'sentBy': value.sentBy,
    };
}

