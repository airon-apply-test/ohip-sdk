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
/**
 * Guest Incident details for the profile.
 * @export
 * @interface IncidentType
 */
export interface IncidentType {
    /**
     * If specified Incident belongs to the Hotel listed, otherwise it is a global Incident.
     * @type {string}
     * @memberof IncidentType
     */
    hotelId?: string;
    /**
     * Code for the Incident.
     * @type {string}
     * @memberof IncidentType
     */
    incidentCode?: string;
    /**
     * Incident Description for display purposes.
     * @type {string}
     * @memberof IncidentType
     */
    description?: string;
    /**
     * Date and Time the Incident was raised.
     * @type {Date}
     * @memberof IncidentType
     */
    incidentOn?: Date;
    /**
     * Status of the Incident.
     * @type {string}
     * @memberof IncidentType
     */
    status?: string;
    /**
     * Incident Priority.
     * @type {string}
     * @memberof IncidentType
     */
    priority?: string;
    /**
     * Source of the incident.
     * @type {string}
     * @memberof IncidentType
     */
    source?: string;
    /**
     * Holds Note for the Incident.
     * @type {string}
     * @memberof IncidentType
     */
    note?: string;
}
/**
 * Check if a given object implements the IncidentType interface.
 */
export declare function instanceOfIncidentType(value: object): boolean;
export declare function IncidentTypeFromJSON(json: any): IncidentType;
export declare function IncidentTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): IncidentType;
export declare function IncidentTypeToJSON(value?: IncidentType | null): any;
